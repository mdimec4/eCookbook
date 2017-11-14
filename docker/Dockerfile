FROM ubuntu:16.04 #FROM armv7/armhf-ubuntu:16.10                                                
# phusion/baseimage:0.9.19 is first release based on ubuntu 16.04.               
#TODO check if there is newer version of the image, that we can use

ENV GOLANG_VERSION 1.9
ENV GOLANG_OS linux
ENV GOLANG_CPU amd64
#ENV NODE_VERSION ${NODE_VERSION:-"7.2.1"}




RUN apt-get update && \
    apt-get install -y --no-install-recommends build-essential && \
    nodejs \
    npm \
    apt-get clean

# dumb init
RUN cd /tmp && \
    wget -P /tmp https://github.com/Yelp/dumb-init/archive/v1.2.0.tar.gz && \
    tar xzf /tmp/v1.2.0.tar.gz -C /tmp/ && \
    cd /tmp/dumb-init-1.2.0 && \
    make && \
    cp dumb-init /sbin && \
    chmod +x /sbin/dumb-init && \
    cd ~ && \
    rm -rf /tmp/dumb-init-1.2.0 /tmp/v1.2.0.tar.gz



# install golang
RUN curl -sSL https://golang.org/dl/go$GOLANG_VERSION.$GOLANG_OS-$GOLANG_CPU.tar.gz | tar -v -C /usr/local -xz
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:$PATH

# install node
# nodejs
#RUN curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/node-v$NODE_VERSION-linux-x64.tar.xz" && \
#   tar xJf "node-v$NODE_VERSION-linux-x64.tar.xz" -C /usr/local --strip-components=1 && \
#    rm "node-v$NODE_VERSION-linux-x64.tar.xz"  && \
#    ln -s /usr/local/bin/node /usr/local/bin/nodejs

RUN apt-get update && \
apt-get install -y --no-install-recommends \
nodejs \
npm &&
apt-get clean && \
ln -s /usr/bin/node /usr/bin/nodejs




# Runs "/usr/bin/dumb-init -- /my/script --with --args"
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
# or if you use --rewrite or other cli flags
# ENTRYPOINT ["dumb-init", "--rewrite", "2:3", "--"]
# CMD ["/my/script", "--with", "--args"]
CMD ["./chef"]




