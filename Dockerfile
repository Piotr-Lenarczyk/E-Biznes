# Use a minimal Debian-based OpenJDK image
FROM openjdk:17-slim

# Set working directory
WORKDIR /app

# Install necessary dependencies, sbt, jq (for json parsing), and procps (for pkill)
RUN apt-get update && apt-get install -y \
    curl \
    gnupg \
    apt-transport-https \
    jq \
    procps \
    && echo "deb https://repo.scala-sbt.org/scalasbt/debian all main" | tee /etc/apt/sources.list.d/sbt.list \
    && echo "deb https://repo.scala-sbt.org/scalasbt/debian /" | tee /etc/apt/sources.list.d/sbt_old.list \
    && curl -sL "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0x2EE0EA64E40A89B84B2DF73499E82A75642AC823" | tee /etc/apt/trusted.gpg.d/sbt.asc \
    && apt-get update && apt-get install -y sbt \
    && rm -rf /var/lib/apt/lists/*

# Copy project files into the container
COPY . .

# Ensure scripts are executable
RUN chmod +x entrypoint.sh test_api.sh

# Set the entrypoint script
ENTRYPOINT ["./entrypoint.sh"]
