md5c is little helper tool for colorizing md5 output.

Run it with a list of files and it prints the md5 sum of the file followed by
the path.

The value in using this over just `md5` is that the hash is colorized and
repeated runs of identical hashes are colored the same. This is used to detect
bands of identical files in a project with lots of repetition.

<img width="928" alt="Example with many bands" src="https://user-images.githubusercontent.com/2027263/57461299-5c151300-723c-11e9-841a-c8ba20860ac0.png">


<img width="822" alt="Example with few bands" src="https://user-images.githubusercontent.com/2027263/57461338-718a3d00-723c-11e9-89f3-cdd10fa10225.png">
