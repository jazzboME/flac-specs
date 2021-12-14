package main

import (
	"fmt"
	"github.com/mewkiz/flac"
	"path/filepath"
)

type summary struct {
	bits     uint8
	rate     uint32
	channels uint8
}

func examine(files []string, detail bool) (string, error) {
	var result string
	var err error

	sumTotal := make(map[summary]int)
	for _, r := range files {
		stream, err := flac.ParseFile((r))
		if err != nil {
			return "", err
		} else {
			defer stream.Close()
			s := stream.Info
			if !detail {
				sumTotal[summary{s.BitsPerSample, s.SampleRate, s.NChannels}] = sumTotal[summary{s.BitsPerSample, s.SampleRate, s.NChannels}] + 1
			} else {
				result += fmt.Sprintf("%s: %v bits/%v kHz/%v channels\n", filepath.Base(r), s.BitsPerSample, s.SampleRate, s.NChannels)
			}
		}

	}

	if detail {
		return result, nil
	} else {
		for spec, numFiles := range sumTotal {
			if numFiles == len(files) {
				result = "All file"
			} else {
				result += fmt.Sprintf("%d file", numFiles)
			}
			if numFiles == 1 && numFiles != len(files) {
				result += " is"
			} else {
				result += "s are"
			}
			result += fmt.Sprintf(" %v bits/%v kHz/%v channels.\n", spec.bits, spec.rate, spec.channels)
		}
	}

	return result, err
}
