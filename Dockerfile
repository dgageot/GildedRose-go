FROM alpine:3.1
MAINTAINER David Gageot <david@gageot.net>

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

RUN apk add --update bash curl git vim go && rm -rf /var/cache/apk/*

RUN go get github.com/dgageot/gildedrose
WORKDIR /gopath/src/github.com/dgageot/gildedrose/

RUN mkdir -p ~/.vim/autoload ~/.vim/bundle && curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
RUN git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go
RUN git clone https://github.com/Shougo/neocomplete.vim.git ~/.vim/bundle/neocomplete.vim
COPY vimrc /root/.vimrc
RUN vim +GoInstallBinaries +qall >/dev/null 2>&1 || true

ENTRYPOINT ["/bin/bash"]
