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
RUN bash -c "source $SDKMAN_DIR/bin/sdkman-init.sh && sdk install kotlin && sdk install gradle"

RUN python3.10 --version
RUN java -version
RUN bash -c "source $SDKMAN_DIR/bin/sdkman-init.sh && kotlin -version && gradle -v"


# Ensure Gradle is in the PATH
ENV PATH="${SDKMAN_DIR}/candidates/gradle/current/bin:${PATH}"
ENV GRADLE_HOME="${SDKMAN_DIR}/candidates/gradle/current"

# Set working directory
WORKDIR /app

# Initialize a Gradle Java application
RUN bash -c "source $SDKMAN_DIR/bin/sdkman-init.sh"

# Configure Gradle build script
RUN echo 'plugins { id "application" }' > /app/build.gradle && \
    echo 'repositories { mavenCentral() }' >> /app/build.gradle && \
    echo 'dependencies { implementation "org.xerial:sqlite-jdbc:3.45.2.0" }' >> /app/build.gradle

RUN cat /app/build.gradle
