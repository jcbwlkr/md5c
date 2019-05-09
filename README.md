md5c is little helper tool for colorizing md5 output.

Run it with a list of files and it prints the md5 sum of the file followed by
the path.

The value in using this over just `md5` is that the hash is colorized and
repeated runs of identical hashes are colored the same. This is used to detect
bands of identical files in a project with lots of repetition.
