FROM ubuntu:16.04 
#FROM armv7/armhf-ubuntu:16.10

ADD . /tmp/repo
WORKDIR /opt/chef

# golang node.js npm
RUN apt-get update && \
    apt-get install -y --no-install-recommends build-essential \
    nodejs \
    npm \
    golang \
    wget \ 
    git && \
    #DEBUG
    ls -lh /usr/bin/node* && \
    # fix node vs. nodejs
    ln -s /usr/bin/nodejs /usr/bin/node && \
    # install dumb init
    cd /tmp && \
    wget -P /tmp https://github.com/Yelp/dumb-init/archive/v1.2.0.tar.gz && \
    tar xzf /tmp/v1.2.0.tar.gz -C /tmp/ && \
    cd /tmp/dumb-init-1.2.0 && \
    make && \
    cp dumb-init /sbin && \
    chmod +x /sbin/dumb-init && \
    cd / && \
    rm -rf /tmp/dumb-init-1.2.0 /tmp/v1.2.0.tar.gz && \
    # build
    cd /tmp/repo && \
    make test && \
    make && \
    mkdir -p /opt/chef && \
    cp -r ./build/* /opt/chef && \
    cd / && \
    rm -r /tmp/repo && \
    # cleanup
    apt-get remove -y build-essential \
    nodejs \
    npm \
    golang \
    wget \
    git && \
    rm /usr/bin/node && \
    apt-get autoremove -y && \
    apt-get clean


# Runs "/usr/bin/dumb-init -- /my/script --with --args"
ENTRYPOINT ["/sbin/dumb-init", "--"]
# or if you use --rewrite or other cli flags
# ENTRYPOINT ["dumb-init", "--rewrite", "2:3", "--"]
# CMD ["/my/script", "--with", "--args"]
CMD ["./chef"]




