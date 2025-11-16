# Maintainer Note

Joshua Gardner (@cellofellow) kindly handed off the maintenance of this project to me, @UncleSp1d3r, so he could focus on other endeavors. I appreciate his generosity and am committed to maintaining this project for the benefit of the community.

## Project Overview

This project is a thin wrapper around Pandora.com's JSON API. It provides a Client struct with various methods that interact with the Pandora JSON API's own methods. Each method returns a struct containing the parsed JSON data and an error. You can find all the responses returned by these methods in the responses subpackage. There is also a requests subpackage, but you generally don't need to interact with it; it gets instantiated by the client methods.

## Short-Term Goals

In line with Cellofellow's original objectives, the short-term goals are to:

- Add proper tests to the project.
- Implement proper error handling.
- Create comprehensive documentation.
- Provide clear examples.
- Include usage examples.

Coming from an enterprise development background, I typically aim to implement CI/CD pipelines and other best practices. However, I plan to keep things simple and focus on the essential components needed to maintain the project smoothly in the foreseeable future as a learning project for others. I warmly welcome contributions from the community and will do my best to review and merge them promptly.

## Project Purpose and Disclaimer

This project serves as a reference implementation and demo for interacting with Pandora's unofficial API. It is intended for educational and research purposes. Users must have valid Pandora account credentials and ensure they have the legal right to access the API while adhering to Pandora's Terms of Service. This project is not affiliated with, endorsed by, or connected to Pandora Media, LLC, or its affiliates.

## Thanks

A big thank you to Joshua Gardner (@cellofellow) for his generosity and dedication to the project. We wouldn't be where we are today without his help. I am also very grateful to the community for their support over the years, and I hope to continue being a good steward of the project.
