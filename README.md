# docker-plugin-auth-debug
Simple docker plugin to debug AuthZ requests.

This plugins solely purpose is to print received authorization requests to std-out.
It simply permits every single request. It is useless by itself, but can help debugging
authorization problems or developing new authorization plugins.
