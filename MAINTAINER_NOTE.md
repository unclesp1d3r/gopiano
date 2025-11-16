# Maintainer Note

Joshua Gardner (cellofellow) was kind enough to hand off the maintenance of this project to me, UncleSp1d3r, so that he could focus on other projects. I'm grateful for his generosity and am committed to maintaining this project for the
benefit of the community.

## Project Overview

This project is a thin wrapper around Pandora.com's JSON API. It provides a Client struct with a myriad of methods
which interact with the Pandora JSON API's own methods. Each method returns a struct of the parsed JSON data and an error.
All of the responses that these methods return can be found in the responses subpackage. There is also a requests subpackage
but mostly you don't need to bother with those; they get instantiated by these client methods.

## Short-Term Goals

In keeping with Cellofellow's original goals, the short-term goals are to:

- Add proper tests to the project.
- Add proper error handling to the project.
- Add proper documentation to the project.
- Add proper examples to the project.
- Add proper usage examples to the project.

Coming from an enterprise development background, my desire is typically to implement CI/CD pipelines and other best practices. However, I intend to keep things simple and just implement the necessary components to keep the project running smoothly for the foreseeable future. I strongly welcome contributions from the community and will do my best to review and merge them in a timely manner.

## Thanks

My sincere thanks to Joshua Gardner (cellofellow) for his generosity and commitment to the project. Without his help, this project would not be where it is today. I am also very grateful to the community for their contributions over the years and hope to continue to be a good steward of the project.
