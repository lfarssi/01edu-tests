for this project its the continue of the last one so go to the previous 
project test and test it than come back here 

be careful the standard banner must be the default



 Default to the standard banner
go run . "Hello"                   
go run . "ASCII Art"              
go run . "1234567890"              
go run . "!@#$%^&*()"              
go run . "Lowercase letters"       
go run . "UPPERCASE TEXT"          

 Explicitly specified banners
go run . "ThinkerToy" thinkertoy   
go run . "Shadow Mode" shadow      
go run . "Hello World!" thinkertoy 
go run . "12345 & Symbols" shadow  

 Edge cases with valid characters
go run . " " standard              
go run . "Multi\nline\ntext" shadow 
go run . "Edge Case: Empty space" 
go run . "LongTextWithoutSpaces" 
go run . "Tab	Text	Example" 
go run . "Boundary~`{|}[]Characters" thinkertoy 
go run . "abcdefghijklmnopqrstuvwxyz" standard 
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" thinkertoy 
go run . "Mix123! ABC" shadow      
go run . "Another Test" standard   
go run . "Hello\nWorld" standard             # Valid: Multi-line text with a newline
go run . "Line1\nLine2\nLine3" thinkertoy    # Valid: Multiple lines
go run . "\nHello" shadow                    # Valid: Starts with a newline
go run . "Hello\n" standard                  # Valid: Ends with a newline
go run . "Middle\nof\nText" thinkertoy       # Valid: Multiple newlines in the middle
go run . "ASCII\nArt" shadow                 # Valid: Text with a newline
go run . "123\n456\n7890"                    # Valid: Numbers with newlines, default banner
go run . "Special\nCharacters\n!@#" standard # Valid: Mixed special characters and newlines
go run . "Single\n" thinkertoy               # Valid: Single word followed by newline
go run . "New\n\nLine" shadow                # Valid: Two consecutive newlines in text
go run . "\n\n" standard                     # Valid: Only newlines
go run . "Test\n\nCase" thinkertoy           # Valid: Mixed text and double newlines
go run . "\n\n\n" shadow                     # Valid: Three newlines in input
go run . "Valid\nText\nHere"                 # Valid: Newlines between words, default banner
go run . "End with newline\n" standard       # Valid: Ends with a newline character
go run . "\nStart with newline" thinkertoy   # Valid: Starts with a newline character
go run . "Multi-line\nInput" shadow          # Valid: Two-line input
go run . "Spaced\n   Text" standard          # Valid: Newline followed by spaces
go run . "Tabs\n\tText" thinkertoy           # Valid: Tab following a newline
go run . "Complex\n\nTest\nCase" shadow      # Valid: Multiple newlines and mixed content


# Missing input
go run .                        # Error: Missing input text

# Too many arguments
go run . "Hello" standard shadow  # Error: Too many arguments
go run . "Test" shadow thinkertoy # Error: Too many arguments

# Invalid banner
go run . "Hello" invalidbanner    # Error: Invalid banner specified
go run . "Test123" 123            # Error: Invalid banner specified


# Out-of-range characters
go run . "Invalidÿ" shadow        # Error: Contains out-of-range characters
go run . "Textë" thinkertoy       # Error: Contains out-of-range characters
go run . "Hello€" standard        # Error: Contains out-of-range characters
go run . "Test💖" shadow          # Error: Emoji or unsupported Unicode
go run . "Text with extra "args""
go run . "Valid" "" 


