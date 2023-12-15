<p align="center">
    <img src="screenshots/0.jpg" width="250" />
    <img src="screenshots/1.jpg" width="250" />
    <img src="screenshots/2.jpg" width="250" />
</p>

## CGI File Upload
>Simple file uploading webpage backend with no clientside JS.

## Usage
This works similarly to cgit or other cgi programs, the executable is called everytime a user makes a request and replies with print statements.
It's pretty rudimentary and it doesn't really have a point.

You can use it with Nginx and a CGI wrapper such as fcgiwrap. \
Make sure the fcgiwrap unix sock file is owned by an user which can write to the StorageDir. Also make sure nginx can access the socket.

You can also use apache2, but I have no experience using it.

## TODO
- [x] Handle existing files (ask to override)
- [x] optionally store login client side, **only store password hash**
- [ ] UI to browse, download and delete files
