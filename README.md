csv
===

A very simple csv pretty printer

Usage
-----

	> csv -h
	Usage of csv:
	  -c=";": seperator char used for parsing
	  -e="": input encoding, e.g. latin9, defaults to UTF-8
	  -o="": output encoding, e.g. latin9, defaults to LC_ALL/LANG or UTF-8
	  -s="|": seperator string used for printing
	
	> csv -e latin1 file.csv
	Röt   | Rotationsachse | 14621001 | 15 | 24
	Mäsk  | Katzenmaske    | 23042    | 17 | 29
	Stäub | Waldfeenstaub  | 2373     | 43 | 81

TODO
----

* Print horizontally (one record per column, not per row)
