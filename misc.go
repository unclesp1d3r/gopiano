package gopiano

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/unclesp1d3r/gopiano/requests"
	"github.com/unclesp1d3r/gopiano/responses"
)

// ExplainTrack retrieves an incomplete list of attributes assigned to a specified song by the
// Music Genome Project.
// Calls API method "track.explainTrack".
func (c *Client) ExplainTrack(ctx context.Context, trackToken string) (*responses.ExplainTrack, error) {
	if err := c.validateUserAuthToken("explaining track"); err != nil {
		return nil, err
	}
	requestData := requests.ExplainTrack{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.ExplainTrack
	err = c.BlowfishCall(ctx, "http://", "track.explainTrack", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MusicSearch searches for music, which can be used to create a new or add seeds to a station.
// Calls API method "music.search".
func (c *Client) MusicSearch(ctx context.Context, searchText string) (*responses.MusicSearch, error) {
	if err := c.validateUserAuthToken("searching music"); err != nil {
		return nil, err
	}
	requestData := requests.MusicSearch{
		SearchText:    searchText,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.MusicSearch
	err = c.BlowfishCall(ctx, "http://", "music.search", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BookmarkAddArtistBookmark bookmarks an artist.
// Argument trackToken is a token of a specific artist.
// Calls API method "bookmark.addArtistBookmark".
func (c *Client) BookmarkAddArtistBookmark(
	ctx context.Context,
	trackToken string,
) (*responses.BookmarkAddArtistBookmark, error) {
	if err := c.validateUserAuthToken("bookmarking artist"); err != nil {
		return nil, err
	}
	requestData := requests.BookmarkAddArtistBookmark{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.BookmarkAddArtistBookmark
	err = c.BlowfishCall(ctx, "http://", "bookmark.addArtistBookmark", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BookmarkAddSongBookmark bookmarks a song.
// Argument trackToken is a token of a specific song.
// Calls API method "bookmark.addSongBookmark".
func (c *Client) BookmarkAddSongBookmark(
	ctx context.Context,
	trackToken string,
) (*responses.BookmarkAddSongBookmark, error) {
	if err := c.validateUserAuthToken("bookmarking song"); err != nil {
		return nil, err
	}
	requestData := requests.BookmarkAddSongBookmark{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.BookmarkAddSongBookmark
	err = c.BlowfishCall(ctx, "http://", "bookmark.addSongBookmark", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
