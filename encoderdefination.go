// add package 
package encoderdefination
import "encoding/hex" 

func EncodedSampleString(s string) (string,int){
	b_string:=[]byte(s)
	encoded_string:=hex.EncodeToString(b_string[:])
	length:=LengthOfString(s)
	return encoded_string,length
}