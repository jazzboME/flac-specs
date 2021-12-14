# flac-specs
Simple cli utility to show bitrate/samples of flac files in a directory

I needed a quick little utility that would look at the flac files in a directory and show me the bitrate, samples and channels of the files. Yes, I probably could have wrapped metaflac with a bash script to do this, but I wanted to try this in go and see how the flac library worked.

There are two command line options:
- `detail` : set to true to show the contents of each file
- `path` : set a specific path, otherwise it assumes `./*.flac`
