module example.com/ashishworkspace/main

go 1.19

replace ashish.com/backend/database => ./backend/

require (
	ashish.com/backend/database v0.0.0-00010101000000-000000000000
	connector v0.0.0-00010101000000-000000000000
)

replace connector => ./connector/
