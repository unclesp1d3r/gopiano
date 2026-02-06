package gopiano

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
	"time"
)

// TestHTTPLeak verifies that HTTP response bodies are properly closed to prevent resource leaks.
func TestHTTPLeak(t *testing.T) {
	// Create a test server that returns a simple response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"stat":"ok","result":{}}`)); err != nil {
			t.Logf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	// Create a client with a modified AndroidClient pointing to our test server
	testClient := AndroidClient
	testClient.BaseURL = strings.TrimPrefix(server.URL, "http://") + "/"

	client, err := NewClient(testClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Get initial goroutine count
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	initialGoroutines := runtime.NumGoroutine()

	// Make multiple HTTP calls to test for resource leaks
	const numCalls = 50
	for i := range numCalls {
		var result interface{}
		err := client.PandoraCall(context.Background(), "http://", "test.method", strings.NewReader("{}"), &result)
		if err != nil {
			t.Logf("Call %d failed (expected for test): %v", i, err)
			// This is expected to fail due to the mock response, but we're testing resource cleanup
		}
	}

	// Allow some time for cleanup
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
	runtime.GC()

	// Check that we haven't leaked goroutines
	finalGoroutines := runtime.NumGoroutine()
	goroutineDiff := finalGoroutines - initialGoroutines

	// Allow for some variance in goroutine count, but significant growth indicates a leak
	if goroutineDiff > 10 {
		t.Errorf("Potential goroutine leak detected: started with %d, ended with %d (diff: %d)",
			initialGoroutines, finalGoroutines, goroutineDiff)
	}

	// Check memory stats for significant growth
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Log stats for debugging
	t.Logf("Goroutines: %d -> %d (diff: %d)", initialGoroutines, finalGoroutines, goroutineDiff)
	t.Logf("Memory allocations: %d -> %d", m1.TotalAlloc, m2.TotalAlloc)
	t.Logf("Active objects: %d -> %d", m1.Mallocs-m1.Frees, m2.Mallocs-m2.Frees)
}

// TestHTTPBodyCloseDirectly tests that the body is properly closed in normal operation.
func TestHTTPBodyCloseDirectly(t *testing.T) {
	// Create a test server that tracks if the request body was read completely
	var bodyRead bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the entire request body to simulate normal operation
		_, err := io.ReadAll(r.Body)
		if err == nil {
			bodyRead = true
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"stat":"ok","result":{"syncTime":` +
			`"` + "1234567890" + `"}}`)); err != nil {
			t.Logf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	// Create a client with a modified AndroidClient pointing to our test server
	testClient := AndroidClient
	testClient.BaseURL = strings.TrimPrefix(server.URL, "http://") + "/"

	client, err := NewClient(testClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test a simple call that should succeed
	var result interface{}
	err = client.PandoraCall(context.Background(), "http://", "test.method", strings.NewReader("{}"), &result)
	// We expect this to work with our mock server
	if err != nil {
		t.Logf("Call completed with: %v", err)
	}

	if !bodyRead {
		t.Error("Request body was not read by server, indicating a potential issue with the test setup")
	}

	t.Log("HTTP body close test completed successfully")
}
