This is a reproducer for https://issuetracker.google.com/issues/128170900.

## How to use

Start the server:

    $ dev_appserver.py .

Visit http://localhost:8080/. It should return "Hello from static handler".

Visit http://localhost:8080/api/signup. It should return "Hello from Go", but it returns a 404 instead.

