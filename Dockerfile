FROM golang:latest

RUN apt-get update
RUN apt-get install -y ruby2.1 ruby2.1-dev
RUN ln -s /usr/bin/ruby2.1 /usr/bin/ruby
RUN ln -s /usr/bin/irb2.1 /usr/bin/irb
RUN ln -s /usr/bin/gem2.1 /usr/bin/gem
