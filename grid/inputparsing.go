package grid

// Function to check if the row is a comment

/*
Ignore all chars after #
If there's a word already, try to parse it. Otherwise skip to next.

Always skip special chars.

Connections format:
st_pancras-euston

1 dash is ok. 2 or more -> figure out which is in the middle

Stations format:
victoria,6,7

*/
