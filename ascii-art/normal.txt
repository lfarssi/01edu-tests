this is test for the project ascii art the normal version 

to verifier your project try to run this test 

this ones are error and it must be handled 
 Missing argument (should produce an error)
go run .

 Multiple arguments (should produce an error)
go run . "test" "test"

 Examples with characters below ASCII 32
go run . "Hello"                     # Backspace (ASCII 8)
go run . "Text with a tab	"          # Tab (ASCII 9)
go run . "Line\nBreak"                # Newline (ASCII 10)

 Examples with characters above ASCII 126
go run . "Helloÿ"                     # 'ÿ' (ASCII 255)
go run . "Testé"                      # 'é' (ASCII 233)
go run . "Invalid√π"                  # √ and π are outside 126
go run . "Symbols¿¡§" 
////////////////////////////
this tests must run without crashing or any error 
/////////////////////////
go run . ""
go run . "123\n\n\n\n"
go run . "   \n   " 
go run . " !\"#$%&'()*+,-./"
go run . "0123456789:;<=>?"
go run . "@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`"
go run . "abcdefghijklmnopqrstuvwxyz{|}~"
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
go run . "0123456789"
go run . "!@#$%^&*()_+-=[]{}|;':\",.<>/?"
go run . "The quick brown fox jumps over the lazy dog!"
go run . "   Leading and trailing spaces    "
go run . "Hello\nWorld"
go run . "Line1\nLine2\n\nLine4"
go run . "12345\n67890\n!@#$%"
go run . "Mixed symbols and numbers: @#$% 12345"
go run . "Punctuation test: .,;:!?\"'`~"
go run . "All uppercase: ABCDEFGHIJKLMNOPQRSTUVWXYZ"
go run . "All lowercase: abcdefghijklmnopqrstuvwxyz"
go run . "Special characters only: ~!@#$%^&*()-_=+[]{}|;:',.<>/?"
go run . "12345 ABCDE abcde"
go run . "This is a test of\nmultiple lines with spaces\nand tabs\tlike this."
go run . "Complete ASCII range from 32 to 126:\n !\"#$%&'()*+,
go run . "This\tis\tweird"

////////////////////
this is new lines test 
///////////////////
go run . "Just a long string\nwith no specific\nstructure\nbut lots of lines"
go run . "123\n\n\n\n"
go run . "   \n   " 
go run . "Hello\nWorld"
go run . "\nStart with a newline"
go run . "End with a newline\n"
go run . "Multiple\nnew\nlines\nin\none\ninput"
go run . "Lines with spaces\n   Line 2 with leading spaces\nLine 3\n"
go run . "ASCII test\n !\"#$%&'()*+,-./0123456789:;<=>?\nABCDEFGHIJKLMNOPQRSTUVWXYZ\nabcdefghijklmnopqrstuvwxyz\n{|}~"
