# go_kfk_client need libsyrdkafka   
build  https://github.com/xenbo/kfkclient

# Check with command pkg-config  
[dongbo@localhost]$ pkg-config --list-all |grep syrdkafka  
syrdkafka              libsyrdkafka - The syKafka C/C++ library  

# run
cd go_kfk_client/example/kfk_example  
go build kfk_example3.go  


