# Could use ubuntu:20.04 image, but we need Go to build the Goku binary
FROM golang:1.18.1 AS goku_builder

# default workdir for this image is /go

# Optional: Create a user, with user dirs?

# Update System
RUN apt-get update && \
    apt-get install -y git && \
    apt-get install -y openssh-client

# Access Goku: Need to install Goku binary (Why is it needed?) There are two ways of doing this:
# 1) Copy the binary from host? But then binary needs to be copied to the host folder where Docker has access
# 2) Clone Goku from Git, but need to have SSH access to the private git repo

# - Second approach: Setup SSH stuff so we can clone Goku private repo 
RUN mkdir ${HOME}/.ssh && \
    touch ${HOME}/.ssh/known_hosts && \
    touch ${HOME}/.ssh/config
RUN ssh-keyscan -H github.com >> ${HOME}/.ssh/known_hosts
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> ${HOME}/.ssh/config

# - Create a symlink of "secret" private SSH key, to the needed directory.
# - - Note: Need to select/pass secret from docker-compose -> build -> secrets, to have it accesible here
RUN ln -s /run/secrets/ssh_private_key ${HOME}/.ssh/id_ed25519 

WORKDIR /go-goku
ARG CACHE_DATE=2021-04-23-v2
RUN --mount=type=secret,id=ssh_private_key git clone -v git@github.com:teejays/goku.git

WORKDIR /go-goku/goku
RUN mkdir bin
RUN go build -o bin/goku ./generator


# STAGE 2
# FROM homebrew/ubuntu20.04 as goku_runner
FROM golang:1.18.1 AS goku_runner
# Need go for go importing file

WORKDIR /go-goku

RUN apt-get update
RUN apt install -y openssl
RUN apt install -y python
RUN apt install -y libpq-dev
RUN apt install -y python3-pip
RUN apt -y install postgresql
RUN apt install -y npm

RUN pip3 install psycopg2-binary
RUN pip3 install yamlfmt

RUN npm install --global prettier

# Install yamltodb/dbtoyaml from Pyrseas.
# - Install from pip `RUN pip3 install Pyrseas` installs a slightly outdated version which does not have the latest commits 
# (e.g. https://github.com/perseas/Pyrseas/commit/8fd62a8d28610fe9a3210a2f787430c034e2d2d2 which solves a yamltodb Postgres 12+ compatibility issue) 
# Hence, we will download from git and build ourselves
RUN git clone https://github.com/perseas/Pyrseas.git
WORKDIR /go-goku/Pyrseas
RUN python3 setup.py install

WORKDIR /go-goku
# COPY Goku binary + code
COPY --from=goku_builder /go-goku/goku /go-goku/goku

# Move goku binary to somewhere in PATH, or add it to path
ENV PATH=/go-goku/goku/bin:${PATH}

WORKDIR /go-goku/app

# RUN make

ENTRYPOINT ["tail", "-f", "/dev/null"]