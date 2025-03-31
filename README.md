# Scala

## Prerequisites
This project can be run on OS which has the following installed:
- <b>Git</b>
- <b>Docker</b>
- For usage #3 only: <b>Play Framework</b>

## Installation
Follow the [demo](https://github.com/Piotr-Lenarczyk/E-Biznes/blob/scala/demos/2025-03-31%2010-35-14.mp4) or do the following:
1. Clone the branch
```
git clone -b scala --single-branch https://github.com/Piotr-Lenarczyk/E-Biznes.git
```
2. Make scripts executable
```
cd E-Biznes/
chmod +x *.sh
```
3. Run the first Docker image (Usage #1)
```
docker build --no-cache -t scala-app .
docker run --rm scala-app
```
This will:
- Prepare the environment
- Start the server
- Run the API endpoints tests defined in [test_api.sh](https://github.com/Piotr-Lenarczyk/E-Biznes/blob/scala/test_api.sh)
- Shut the server down
4. Run the second Docker image (Usage #2)
```
docker build --no-cache -f Dockerfile.ngrok -t play-scala-ngrok .
docker run --rm -e NGROK_AUTH_TOKEN=$NGROK_AUTH_TOKEN play-scala-ngrok
```
Replace `$NGROK_AUTH_TOKEN` with your own authorization token<br>
You might need to run the second command twice<br>
Make sure to first <i>run point 3</i> since this Docker image depends on it<br>
This will:
- Prepare the environment
- Start the server
- Expose server's port
- Setup an ngrok tunnel<br>

The server will now be available from public URL shown in the console<br>
This container does not have graceful shutdown mechanisms and when forced to stop, will not release the tunnel. It needs to be done manually<br>
5. Alternatively, server can be run locally according to needs (Usage #3)
```
sbt run
```
