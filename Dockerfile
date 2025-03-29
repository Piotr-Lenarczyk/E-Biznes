FROM ubuntu:24.04
SHELL ["/bin/bash", "-c"]

# Install dependencies
RUN apt update
RUN apt install -y software-properties-common curl unzip zip
RUN add-apt-repository ppa:deadsnakes/ppa -y
RUN apt update
RUN apt install -y python3.10 openjdk-8-jdk
RUN apt clean

# Install SDKMAN and make sure it's available in all future shells
RUN curl -s "https://get.sdkman.io" | bash && \
    echo "source /root/.sdkman/bin/sdkman-init.sh" >> /root/.bashrc

# Set environment variables
ENV SDKMAN_DIR="/root/.sdkman"
ENV PATH="${SDKMAN_DIR}/bin:${PATH}"

# Install Kotlin and Gradle
RUN bash -c "source $SDKMAN_DIR/bin/sdkman-init.sh && sdk install kotlin"
RUN python3.10 --version
RUN java -version
RUN bash -c "source $SDKMAN_DIR/bin/sdkman-init.sh && kotlin -version"
