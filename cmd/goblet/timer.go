package main

import "time"

type FrameTimer struct {
	frameDuration time.Duration
	startTime     time.Time
	lastFrame     int64
}

func NewFrameTimer(fps int) *FrameTimer {
	return &FrameTimer{
		frameDuration: time.Second / time.Duration(fps),
		startTime:     time.Now(),
		lastFrame:     0,
	}
}

func (t *FrameTimer) Delta() float64 {
	return t.frameDuration.Seconds()
}

func (t *FrameTimer) Reset() {
	t.startTime = time.Now()
	t.lastFrame = 0
}

func (t *FrameTimer) Cutoff() int {
	frame := int64(time.Since(t.startTime) / t.frameDuration)
	framesPassed := int(frame - t.lastFrame)
	t.lastFrame = frame

	return framesPassed
}

func (t *FrameTimer) Relax() {
	nextFrameTime := t.startTime.Add(
		time.Duration(t.lastFrame+1) * t.frameDuration)
	time.Sleep(time.Until(nextFrameTime))
}
