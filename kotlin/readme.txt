
sudo snap install --classic kotlin
sudo apt install openjdk-8-jre-headless 

cat hello.kt

fun main(args: Array<String>) {
    println("Hello, World!")
}

kotlinc hello.kt -include-runtime -d hello.jar

java -jar hello.jar 

