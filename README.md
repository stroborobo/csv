csv
===

A very simple csv formatter


Usage
-----

	> csv -h
	Usage of csv:
	  -c=";": seperator char used for parsing
	  -e="": input encoding, e.g. latin9, defaults to UTF-8
	  -s="|": seperator string used for printing
	
	> csv -e latin1 file.csv
	import001 | ImpXD001 ä | 14621001 | 15 | 24
	import002 | ImpXD002 ö | 14621002 | 17 | 29
	import003 | ImpXD003 ü | 14621003 | 43 | 81
