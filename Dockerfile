FROM scratch

ADD bin/api.linux /
CMD ["bin/api.linux"]

EXPOSE 3000