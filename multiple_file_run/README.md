### How to run multiple file from one file 
* create the main file
* call another function locate in same the same directory of different go file
* build the Binary File using the following command
    ```bash
  $ go build -o <BinaryFileName> .
    ```
  if you want to specify the file name run the following command
  ```
  $ go build -o <BinaryFileName> <fileName1.go> <fileName2.go>
  ```
* Finally to run the Binary file 
  ```bash
   $ ./BinaryFile
  ```
        