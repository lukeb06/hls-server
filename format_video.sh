# !/bin/bash

input_file=$1
output_name=$2
chunk_size=$3

# ex.
# ffmpeg -i {file} -codec: copy -start_number 0 -hls_time {chunk_size} -hls_list_size 0 -f hls output.m3u8

ffmpeg -i $input_file -codec: copy -start_number 0 -hls_time $chunk_size -hls_list_size 0 -f hls hls/$output_name.m3u8

# Example Usage:
# ./format_video.sh /Users/lukebarrier/Downloads/test.mp4 test 10
