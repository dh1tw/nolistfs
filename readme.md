## TL;DR

NoListFileSystem is a custom filesystem implementation. It wraps around a "base" file system. It is mainly used with `http.FileServer` so that 404 status code is returned instead of a directory listing.

Instead of copy & pasting this code in all of my projects, I prefer to store it in a dedicated package.

## Example

Please check the example folder for a basic example which: 
- uses [go embed](https://golang.org/pkg/embed/) to embed all static assets
- shows how to use NoListFileSystem
- runs an HTTP server to demonstrate the behaviour

## Credits

- Alex Edwards on [How to disable http.FileServer Directory Listings](https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings#using-middleware)
