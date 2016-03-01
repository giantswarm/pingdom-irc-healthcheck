FROM scratch
MAINTAINER Joseph Salisbury <joseph@giantswarm.io>

ADD ./pingdom-irc-healthcheck /pingdom-irc-healthcheck

EXPOSE 8000

ENTRYPOINT ["/pingdom-irc-healthcheck"]