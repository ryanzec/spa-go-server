# Single Page Application Server

This is a very simple HTTP server designed to be able to serve Single Page Application.  This came out of frustration of trying to work with python's SimpleHTTPServer to server a ReactJS application.  It would work for the root url however if I navigated to a pathed url like `http://localhost:3000/some-page` and tried to reload from the URL, I would get a 404, I always had to go back to the root to reload the page.  I also wanted a simple project to test out the Go language.
