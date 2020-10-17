#ifndef c_header_h
#define c_header_h

#include <stdlib.h>
#include <stdint.h>
#if __wasm32__

__attribute__((__import_name__("fetch_code")))
extern int c_fetchCode(int input);

#endif
#endif /* c_header_h */
