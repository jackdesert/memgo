Memgo
=====

A Memcached Replacement, built as a learning experience.


POST /
------

Edit the file `sample-post`, then call:

  curl -X POST -d @sample-post localhost:3000 --header "Content-Type:application/json"


GET /:key
---------

  curl localhost:3000/some-key


