#!/bin/bash

# Create or clear the interface.go file
output_file="interface.go"
echo "// Package zap defines the interface a ZAP client should implement" > $output_file
echo "package zap" >> $output_file
echo "" >> $output_file
echo "// Interface defines the interface a ZAP client should implement" >> $output_file
echo "type Interface interface {" >> $output_file

# Find all _generated.go files
generated_files=$(find . -name "*_generated.go")

# Loop through each _generated.go file and extract the struct and function
for file in $generated_files; do
    struct_name=$(grep -o 'type [^ ]* struct' "$file" | awk '{print $2}')
    if [ -n "$struct_name" ]; then
        echo "    $struct_name() *$struct_name" >> $output_file
    fi
done

echo "}" >> $output_file
echo "" >> $output_file

# Loop through each _generated.go file again to extract and write function implementations
for file in $generated_files; do
    struct_name=$(grep -o 'type [^ ]* struct' "$file" | awk '{print $2}')
    if [ -n "$struct_name" ]; then
        echo "// $struct_name() returns a $struct_name client" >> $output_file
        echo "func (c *Client) $struct_name() *$struct_name {" >> $output_file
        echo "    return &$struct_name{c}" >> $output_file
        echo "}" >> $output_file
        echo "" >> $output_file
    fi
done

echo "interface.go file has been generated."
