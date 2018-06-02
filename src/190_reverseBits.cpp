#include <iostream>
#include <stdint.h>

class Solution {
public:
	uint32_t reverseBits(uint32_t n) {
	    uint32_t result = 0;

	    for(int i = 0; i < 32; i++){
	    	result = result<<1;
	        result |= (n&1);
	        n= n>>1;
	    }
	    
	    return result;
	}
};

uint32_t reverseBits(uint32_t n) {
    uint32_t result = 0;

    for(int i = 0; i < 32; i++){
    	result = result<<1;
        result |= (n&1);
        n= n>>1;
    }
    
    return result;
}

int main(int argc,char *argv[])
{
    std::cout << reverseBits(2) << std::endl;
    return(0);
}