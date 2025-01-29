# bigdata1

This is assignment 1 for bigdata platforms


install kafka :

Kafka requires Java 8 or newer. You can install OpenJDK on your system:
# ubuntu/linux
sudo apt install openjdk-11-jdk 
# mac
brew install openjdk@11

Go to the Apache Kafka download page and download the latest stable release of Kafka.

tar -xvf kafka-3.9.0-src.tgz
 cd kafka-3.9.0-src    

for the first time run this: 
./gradlew jar -PscalaVersion=2.13.14

start zookeeper 
bin/zookeeper-server-start.sh config/zookeeper.properties

to start kafka in another terminal
bin/kafka-server-start.sh config/server.properties

to create topic
bin/kafka-topics.sh --create --topic test-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1











