package audit

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

func Info(ctx context.Context, topic, msg string) {
	log.Warningf(ctx, "'%v': %v", topic, msg)
}

func Warning(ctx context.Context, topic, msg string) {
	log.Warningf(ctx, "'%v': %v", topic, msg)
}

func Critical(ctx context.Context, topic, msg string) {
	log.Criticalf(ctx, "'%v': %v", topic, msg)
}

func Debug(ctx context.Context, topic, msg string) {
	log.Debugf(ctx, "'%v': %v", topic, msg)
}

func Error(ctx context.Context, topic string, err error) {
	log.Errorf(ctx, "'%v': %v", topic, err)
}
