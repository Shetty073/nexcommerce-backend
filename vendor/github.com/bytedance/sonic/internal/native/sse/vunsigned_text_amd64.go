// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package sse

var _text_vunsigned = []byte{
	// .p2align 4, 0x90
	// _vunsigned
	0x55, // pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000001 movq         %rsp, %rbp
	0x41, 0x56, //0x00000004 pushq        %r14
	0x53, //0x00000006 pushq        %rbx
	0x49, 0x89, 0xd0, //0x00000007 movq         %rdx, %r8
	0x48, 0x8b, 0x0e, //0x0000000a movq         (%rsi), %rcx
	0x4c, 0x8b, 0x0f, //0x0000000d movq         (%rdi), %r9
	0x4c, 0x8b, 0x77, 0x08, //0x00000010 movq         $8(%rdi), %r14
	0x48, 0xc7, 0x02, 0x09, 0x00, 0x00, 0x00, //0x00000014 movq         $9, (%rdx)
	0x0f, 0x57, 0xc0, //0x0000001b xorps        %xmm0, %xmm0
	0x0f, 0x11, 0x42, 0x08, //0x0000001e movups       %xmm0, $8(%rdx)
	0x48, 0x8b, 0x06, //0x00000022 movq         (%rsi), %rax
	0x48, 0x89, 0x42, 0x18, //0x00000025 movq         %rax, $24(%rdx)
	0x4c, 0x39, 0xf1, //0x00000029 cmpq         %r14, %rcx
	0x0f, 0x83, 0x1b, 0x00, 0x00, 0x00, //0x0000002c jae          LBB0_1
	0x41, 0x8a, 0x04, 0x09, //0x00000032 movb         (%r9,%rcx), %al
	0x3c, 0x2d, //0x00000036 cmpb         $45, %al
	0x0f, 0x85, 0x1e, 0x00, 0x00, 0x00, //0x00000038 jne          LBB0_4
	//0x0000003e LBB0_3
	0x48, 0x89, 0x0e, //0x0000003e movq         %rcx, (%rsi)
	0x49, 0xc7, 0x00, 0xfa, 0xff, 0xff, 0xff, //0x00000041 movq         $-6, (%r8)
	0x5b, //0x00000048 popq         %rbx
	0x41, 0x5e, //0x00000049 popq         %r14
	0x5d, //0x0000004b popq         %rbp
	0xc3, //0x0000004c retq         
	//0x0000004d LBB0_1
	0x4c, 0x89, 0x36, //0x0000004d movq         %r14, (%rsi)
	0x49, 0xc7, 0x00, 0xff, 0xff, 0xff, 0xff, //0x00000050 movq         $-1, (%r8)
	0x5b, //0x00000057 popq         %rbx
	0x41, 0x5e, //0x00000058 popq         %r14
	0x5d, //0x0000005a popq         %rbp
	0xc3, //0x0000005b retq         
	//0x0000005c LBB0_4
	0x8d, 0x50, 0xc6, //0x0000005c leal         $-58(%rax), %edx
	0x80, 0xfa, 0xf5, //0x0000005f cmpb         $-11, %dl
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x00000062 ja           LBB0_6
	0x48, 0x89, 0x0e, //0x00000068 movq         %rcx, (%rsi)
	0x49, 0xc7, 0x00, 0xfe, 0xff, 0xff, 0xff, //0x0000006b movq         $-2, (%r8)
	0x5b, //0x00000072 popq         %rbx
	0x41, 0x5e, //0x00000073 popq         %r14
	0x5d, //0x00000075 popq         %rbp
	0xc3, //0x00000076 retq         
	//0x00000077 LBB0_6
	0x3c, 0x30, //0x00000077 cmpb         $48, %al
	0x0f, 0x85, 0x26, 0x00, 0x00, 0x00, //0x00000079 jne          LBB0_10
	0x41, 0x8a, 0x44, 0x09, 0x01, //0x0000007f movb         $1(%r9,%rcx), %al
	0x04, 0xd2, //0x00000084 addb         $-46, %al
	0x3c, 0x37, //0x00000086 cmpb         $55, %al
	0x0f, 0x87, 0xc7, 0x00, 0x00, 0x00, //0x00000088 ja           LBB0_9
	0x0f, 0xb6, 0xc0, //0x0000008e movzbl       %al, %eax
	0x48, 0xba, 0x01, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80, 0x00, //0x00000091 movabsq      $36028797027352577, %rdx
	0x48, 0x0f, 0xa3, 0xc2, //0x0000009b btq          %rax, %rdx
	0x0f, 0x83, 0xb0, 0x00, 0x00, 0x00, //0x0000009f jae          LBB0_9
	//0x000000a5 LBB0_10
	0x49, 0x39, 0xce, //0x000000a5 cmpq         %rcx, %r14
	0x49, 0x89, 0xca, //0x000000a8 movq         %rcx, %r10
	0x4d, 0x0f, 0x47, 0xd6, //0x000000ab cmovaq       %r14, %r10
	0x31, 0xc0, //0x000000af xorl         %eax, %eax
	0x41, 0xbb, 0x0a, 0x00, 0x00, 0x00, //0x000000b1 movl         $10, %r11d
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000000b7 .p2align 4, 0x90
	//0x000000c0 LBB0_11
	0x4c, 0x39, 0xf1, //0x000000c0 cmpq         %r14, %rcx
	0x0f, 0x83, 0x80, 0x00, 0x00, 0x00, //0x000000c3 jae          LBB0_22
	0x41, 0x0f, 0xbe, 0x1c, 0x09, //0x000000c9 movsbl       (%r9,%rcx), %ebx
	0x8d, 0x53, 0xd0, //0x000000ce leal         $-48(%rbx), %edx
	0x80, 0xfa, 0x09, //0x000000d1 cmpb         $9, %dl
	0x0f, 0x87, 0x44, 0x00, 0x00, 0x00, //0x000000d4 ja           LBB0_17
	0x49, 0xf7, 0xe3, //0x000000da mulq         %r11
	0x0f, 0x80, 0x28, 0x00, 0x00, 0x00, //0x000000dd jo           LBB0_16
	0x48, 0x83, 0xc1, 0x01, //0x000000e3 addq         $1, %rcx
	0x83, 0xc3, 0xd0, //0x000000e7 addl         $-48, %ebx
	0x31, 0xff, //0x000000ea xorl         %edi, %edi
	0x48, 0x01, 0xd8, //0x000000ec addq         %rbx, %rax
	0x40, 0x0f, 0x92, 0xc7, //0x000000ef setb         %dil
	0x48, 0x89, 0xfa, //0x000000f3 movq         %rdi, %rdx
	0x48, 0xf7, 0xda, //0x000000f6 negq         %rdx
	0x48, 0x31, 0xd7, //0x000000f9 xorq         %rdx, %rdi
	0x0f, 0x85, 0x09, 0x00, 0x00, 0x00, //0x000000fc jne          LBB0_16
	0x48, 0x85, 0xd2, //0x00000102 testq        %rdx, %rdx
	0x0f, 0x89, 0xb5, 0xff, 0xff, 0xff, //0x00000105 jns          LBB0_11
	//0x0000010b LBB0_16
	0x48, 0x83, 0xc1, 0xff, //0x0000010b addq         $-1, %rcx
	0x48, 0x89, 0x0e, //0x0000010f movq         %rcx, (%rsi)
	0x49, 0xc7, 0x00, 0xfb, 0xff, 0xff, 0xff, //0x00000112 movq         $-5, (%r8)
	0x5b, //0x00000119 popq         %rbx
	0x41, 0x5e, //0x0000011a popq         %r14
	0x5d, //0x0000011c popq         %rbp
	0xc3, //0x0000011d retq         
	//0x0000011e LBB0_17
	0x4c, 0x39, 0xf1, //0x0000011e cmpq         %r14, %rcx
	0x0f, 0x83, 0x1f, 0x00, 0x00, 0x00, //0x00000121 jae          LBB0_21
	0x41, 0x8a, 0x14, 0x09, //0x00000127 movb         (%r9,%rcx), %dl
	0x80, 0xfa, 0x2e, //0x0000012b cmpb         $46, %dl
	0x0f, 0x84, 0x0a, 0xff, 0xff, 0xff, //0x0000012e je           LBB0_3
	0x80, 0xfa, 0x45, //0x00000134 cmpb         $69, %dl
	0x0f, 0x84, 0x01, 0xff, 0xff, 0xff, //0x00000137 je           LBB0_3
	0x80, 0xfa, 0x65, //0x0000013d cmpb         $101, %dl
	0x0f, 0x84, 0xf8, 0xfe, 0xff, 0xff, //0x00000140 je           LBB0_3
	//0x00000146 LBB0_21
	0x49, 0x89, 0xca, //0x00000146 movq         %rcx, %r10
	//0x00000149 LBB0_22
	0x4c, 0x89, 0x16, //0x00000149 movq         %r10, (%rsi)
	0x49, 0x89, 0x40, 0x10, //0x0000014c movq         %rax, $16(%r8)
	0x5b, //0x00000150 popq         %rbx
	0x41, 0x5e, //0x00000151 popq         %r14
	0x5d, //0x00000153 popq         %rbp
	0xc3, //0x00000154 retq         
	//0x00000155 LBB0_9
	0x48, 0x83, 0xc1, 0x01, //0x00000155 addq         $1, %rcx
	0x48, 0x89, 0x0e, //0x00000159 movq         %rcx, (%rsi)
	0x5b, //0x0000015c popq         %rbx
	0x41, 0x5e, //0x0000015d popq         %r14
	0x5d, //0x0000015f popq         %rbp
	0xc3, //0x00000160 retq         
	0x00, 0x00, 0x00, //0x00000161 .p2align 2, 0x00
	//0x00000164 _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x00000164 .long 2
}
 
