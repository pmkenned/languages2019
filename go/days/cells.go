package main

import (
	"fmt"
	"strings"
	"strconv"
)

type claimT struct {
	row int
	col int
	width int
	height int
}

func main() {

	claims_raw := `#1 @ 45,64: 22x22
#2 @ 291,191: 12x19
#3 @ 497,49: 24x27
#4 @ 694,889: 28x28
#5 @ 942,231: 15x24
#6 @ 463,561: 19x10
#7 @ 756,392: 27x19
#8 @ 612,291: 21x17
#9 @ 86,585: 12x29
#10 @ 230,801: 27x24
#11 @ 823,833: 28x13
#12 @ 204,231: 28x25
#13 @ 2,434: 28x24
#14 @ 495,349: 15x23
#15 @ 867,820: 29x15
#16 @ 962,23: 19x16
#17 @ 858,412: 19x11
#18 @ 844,806: 23x23
#19 @ 57,461: 14x23
#20 @ 870,664: 21x14
#21 @ 140,653: 17x13
#22 @ 527,886: 13x23
#23 @ 960,551: 25x15
#24 @ 902,201: 29x22
#25 @ 43,343: 19x15
#26 @ 304,311: 10x29
#27 @ 965,403: 28x27
#28 @ 396,928: 10x10
#29 @ 309,201: 10x13
#30 @ 116,268: 10x15
#31 @ 937,429: 26x23
#32 @ 936,488: 25x20
#33 @ 139,735: 12x16
#34 @ 694,523: 24x28
#35 @ 190,574: 29x24
#36 @ 96,457: 17x16
#37 @ 125,516: 15x27
#38 @ 770,701: 16x10
#39 @ 233,834: 18x12
#40 @ 957,746: 19x20
#41 @ 271,508: 15x20
#42 @ 788,790: 28x13
#43 @ 115,972: 28x27
#44 @ 561,140: 21x26
#45 @ 326,365: 26x24
#46 @ 867,489: 10x24
#47 @ 662,648: 22x18
#48 @ 913,592: 14x18
#49 @ 409,816: 11x18
#50 @ 310,86: 11x27
#51 @ 956,316: 26x11
#52 @ 219,987: 13x10
#53 @ 884,334: 17x18
#54 @ 732,557: 21x23
#55 @ 465,890: 29x11
#56 @ 346,376: 28x17
#57 @ 740,736: 12x11
#58 @ 916,484: 15x23
#59 @ 570,698: 28x15
#60 @ 530,868: 17x28
#61 @ 196,427: 19x26
#62 @ 48,912: 27x27
#63 @ 734,807: 16x24
#64 @ 709,190: 13x27
#65 @ 627,620: 12x17
#66 @ 162,446: 22x18
#67 @ 16,307: 20x16
#68 @ 156,655: 29x28
#69 @ 682,488: 14x24
#70 @ 151,970: 17x19
#71 @ 713,852: 27x12
#72 @ 763,85: 28x16
#73 @ 487,584: 12x10
#74 @ 27,731: 24x27
#75 @ 454,45: 14x10
#76 @ 760,347: 26x16
#77 @ 278,453: 14x18
#78 @ 553,40: 10x27
#79 @ 459,390: 17x23
#80 @ 621,630: 22x25
#81 @ 715,507: 20x14
#82 @ 818,646: 19x16
#83 @ 593,564: 26x28
#84 @ 759,685: 20x10
#85 @ 707,867: 23x22
#86 @ 65,25: 22x27
#87 @ 75,74: 14x10
#88 @ 180,27: 28x27
#89 @ 725,511: 18x29
#90 @ 145,768: 21x18
#91 @ 407,744: 25x20
#92 @ 135,434: 28x18
#93 @ 31,861: 21x11
#94 @ 429,848: 23x29
#95 @ 273,171: 15x26
#96 @ 778,344: 12x28
#97 @ 516,911: 29x26
#98 @ 622,7: 24x24
#99 @ 522,518: 26x26
#100 @ 46,242: 25x22
#101 @ 343,392: 28x17
#102 @ 3,816: 24x20
#103 @ 55,492: 28x28
#104 @ 377,100: 23x11
#105 @ 695,483: 12x10
#106 @ 409,854: 19x19
#107 @ 387,79: 23x24
#108 @ 871,898: 29x10
#109 @ 687,351: 23x19
#110 @ 719,746: 19x29
#111 @ 681,894: 10x23
#112 @ 362,833: 26x24
#113 @ 524,516: 18x19
#114 @ 257,161: 27x17
#115 @ 970,563: 23x23
#116 @ 861,727: 29x17
#117 @ 339,372: 19x18
#118 @ 458,824: 29x12
#119 @ 690,961: 5x3
#120 @ 631,182: 11x24
#121 @ 859,350: 13x27
#122 @ 4,494: 17x23
#123 @ 266,263: 12x24
#124 @ 521,732: 14x18
#125 @ 656,729: 22x18
#126 @ 756,347: 19x25
#127 @ 517,595: 21x20
#128 @ 414,452: 16x23
#129 @ 170,233: 24x10
#130 @ 721,205: 16x16
#131 @ 899,786: 28x16
#132 @ 596,362: 10x23
#133 @ 841,709: 22x22
#134 @ 637,561: 19x15
#135 @ 220,654: 22x23
#136 @ 235,187: 6x13
#137 @ 0,742: 17x29
#138 @ 740,143: 28x21
#139 @ 45,547: 29x16
#140 @ 730,712: 14x28
#141 @ 217,944: 16x23
#142 @ 214,372: 12x11
#143 @ 921,609: 11x18
#144 @ 72,764: 29x19
#145 @ 143,643: 26x18
#146 @ 895,965: 18x14
#147 @ 743,447: 17x15
#148 @ 948,643: 25x28
#149 @ 214,774: 22x18
#150 @ 16,16: 18x20
#151 @ 20,473: 23x28
#152 @ 456,702: 11x24
#153 @ 974,294: 26x16
#154 @ 406,650: 18x16
#155 @ 453,847: 22x21
#156 @ 802,432: 12x10
#157 @ 645,915: 22x18
#158 @ 561,931: 18x22
#159 @ 291,742: 25x20
#160 @ 87,23: 16x25
#161 @ 161,0: 17x16
#162 @ 491,414: 20x22
#163 @ 69,678: 24x29
#164 @ 359,362: 24x15
#165 @ 777,265: 29x14
#166 @ 984,660: 12x27
#167 @ 864,802: 17x21
#168 @ 498,541: 11x23
#169 @ 954,149: 28x25
#170 @ 211,931: 20x24
#171 @ 177,547: 20x17
#172 @ 852,380: 15x18
#173 @ 584,175: 20x28
#174 @ 858,353: 29x15
#175 @ 707,312: 29x29
#176 @ 654,226: 11x28
#177 @ 250,362: 13x29
#178 @ 228,373: 27x22
#179 @ 320,71: 22x18
#180 @ 262,936: 27x12
#181 @ 913,478: 25x19
#182 @ 139,65: 13x29
#183 @ 405,781: 10x12
#184 @ 626,194: 19x19
#185 @ 729,105: 25x26
#186 @ 50,770: 13x24
#187 @ 534,637: 10x28
#188 @ 17,441: 10x20
#189 @ 114,253: 20x17
#190 @ 55,595: 20x10
#191 @ 216,225: 27x16
#192 @ 11,531: 13x11
#193 @ 849,545: 15x14
#194 @ 58,650: 11x11
#195 @ 688,789: 13x13
#196 @ 690,354: 23x11
#197 @ 839,816: 22x23
#198 @ 890,583: 26x26
#199 @ 686,426: 22x16
#200 @ 410,840: 24x24
#201 @ 855,308: 13x29
#202 @ 155,374: 17x27
#203 @ 74,152: 22x12
#204 @ 693,546: 25x12
#205 @ 412,286: 28x14
#206 @ 630,565: 13x15
#207 @ 526,494: 28x22
#208 @ 711,900: 21x27
#209 @ 737,557: 24x12
#210 @ 524,603: 10x3
#211 @ 671,696: 23x16
#212 @ 928,650: 25x29
#213 @ 743,223: 29x11
#214 @ 383,234: 26x19
#215 @ 95,450: 10x10
#216 @ 193,367: 10x15
#217 @ 587,303: 12x14
#218 @ 367,119: 11x28
#219 @ 354,783: 27x13
#220 @ 333,381: 25x20
#221 @ 497,362: 22x13
#222 @ 520,973: 10x27
#223 @ 38,868: 10x12
#224 @ 839,814: 19x26
#225 @ 917,920: 14x20
#226 @ 668,272: 26x17
#227 @ 122,505: 13x27
#228 @ 379,251: 7x16
#229 @ 417,482: 29x15
#230 @ 791,126: 19x19
#231 @ 125,728: 20x12
#232 @ 203,326: 11x19
#233 @ 596,571: 28x17
#234 @ 96,10: 10x24
#235 @ 241,868: 27x14
#236 @ 921,216: 22x29
#237 @ 355,386: 10x24
#238 @ 804,689: 28x16
#239 @ 447,38: 16x24
#240 @ 854,572: 28x26
#241 @ 672,465: 29x26
#242 @ 178,287: 22x23
#243 @ 815,830: 15x14
#244 @ 13,111: 23x22
#245 @ 78,256: 16x12
#246 @ 631,195: 10x19
#247 @ 588,219: 14x10
#248 @ 446,21: 16x25
#249 @ 456,635: 14x13
#250 @ 543,55: 14x27
#251 @ 72,658: 24x12
#252 @ 206,72: 22x10
#253 @ 418,394: 14x25
#254 @ 213,26: 22x20
#255 @ 507,543: 23x14
#256 @ 455,400: 21x12
#257 @ 71,907: 18x27
#258 @ 670,185: 24x25
#259 @ 251,560: 13x16
#260 @ 387,973: 16x12
#261 @ 910,601: 12x10
#262 @ 341,792: 18x26
#263 @ 932,800: 24x25
#264 @ 923,939: 21x23
#265 @ 314,88: 3x21
#266 @ 100,438: 12x10
#267 @ 329,628: 26x11
#268 @ 941,497: 22x22
#269 @ 100,342: 18x13
#270 @ 762,65: 13x22
#271 @ 91,594: 11x16
#272 @ 716,12: 29x25
#273 @ 755,640: 13x27
#274 @ 299,963: 18x29
#275 @ 264,333: 22x23
#276 @ 46,660: 19x25
#277 @ 683,623: 15x26
#278 @ 362,726: 23x24
#279 @ 665,862: 25x15
#280 @ 778,15: 19x11
#281 @ 302,775: 21x29
#282 @ 180,361: 14x16
#283 @ 627,32: 11x27
#284 @ 440,549: 27x14
#285 @ 477,522: 26x27
#286 @ 112,322: 15x20
#287 @ 913,540: 17x10
#288 @ 213,730: 29x12
#289 @ 878,471: 26x19
#290 @ 435,338: 17x20
#291 @ 930,502: 14x19
#292 @ 83,602: 18x15
#293 @ 311,79: 13x20
#294 @ 941,478: 19x25
#295 @ 82,935: 19x24
#296 @ 620,339: 20x29
#297 @ 526,75: 15x13
#298 @ 7,632: 25x29
#299 @ 579,110: 26x22
#300 @ 267,881: 27x21
#301 @ 784,871: 21x10
#302 @ 328,480: 28x10
#303 @ 93,950: 18x11
#304 @ 914,241: 28x26
#305 @ 789,3: 19x16
#306 @ 420,744: 22x15
#307 @ 148,342: 19x12
#308 @ 661,725: 21x22
#309 @ 240,502: 15x10
#310 @ 80,205: 29x28
#311 @ 402,286: 14x21
#312 @ 497,526: 28x17
#313 @ 750,436: 24x19
#314 @ 869,571: 19x14
#315 @ 975,561: 10x18
#316 @ 45,518: 13x15
#317 @ 776,859: 27x25
#318 @ 622,607: 25x21
#319 @ 69,48: 29x19
#320 @ 751,479: 22x10
#321 @ 561,188: 27x28
#322 @ 834,962: 29x23
#323 @ 670,142: 27x21
#324 @ 39,357: 10x29
#325 @ 920,292: 29x23
#326 @ 658,863: 26x22
#327 @ 941,431: 29x13
#328 @ 245,620: 18x23
#329 @ 81,254: 20x28
#330 @ 417,626: 28x26
#331 @ 173,358: 18x28
#332 @ 688,348: 12x14
#333 @ 779,344: 22x19
#334 @ 670,678: 18x25
#335 @ 831,793: 17x27
#336 @ 264,229: 14x25
#337 @ 483,189: 23x21
#338 @ 48,222: 28x21
#339 @ 853,34: 26x12
#340 @ 32,817: 25x22
#341 @ 961,626: 22x16
#342 @ 920,618: 10x10
#343 @ 882,753: 12x25
#344 @ 25,802: 26x28
#345 @ 254,427: 25x14
#346 @ 691,550: 16x29
#347 @ 566,347: 12x28
#348 @ 368,22: 21x25
#349 @ 457,55: 27x24
#350 @ 170,60: 27x10
#351 @ 835,781: 22x20
#352 @ 496,873: 19x22
#353 @ 320,839: 14x19
#354 @ 858,698: 18x15
#355 @ 705,478: 17x13
#356 @ 253,336: 16x14
#357 @ 421,597: 29x14
#358 @ 387,437: 25x28
#359 @ 337,239: 28x19
#360 @ 8,824: 14x27
#361 @ 800,457: 18x28
#362 @ 855,283: 29x17
#363 @ 729,554: 13x18
#364 @ 340,438: 29x23
#365 @ 649,917: 23x19
#366 @ 146,607: 12x22
#367 @ 186,590: 10x19
#368 @ 913,877: 17x17
#369 @ 581,590: 17x20
#370 @ 527,472: 19x21
#371 @ 227,6: 27x29
#372 @ 415,697: 25x10
#373 @ 838,410: 24x22
#374 @ 866,943: 10x10
#375 @ 171,425: 22x16
#376 @ 882,757: 23x10
#377 @ 108,126: 16x17
#378 @ 604,920: 28x20
#379 @ 100,651: 22x17
#380 @ 883,261: 22x22
#381 @ 3,856: 13x14
#382 @ 159,424: 21x24
#383 @ 207,537: 15x13
#384 @ 147,703: 4x4
#385 @ 389,76: 24x11
#386 @ 745,739: 25x10
#387 @ 539,819: 10x26
#388 @ 916,227: 10x21
#389 @ 844,314: 17x27
#390 @ 480,589: 21x13
#391 @ 168,684: 17x10
#392 @ 320,334: 22x25
#393 @ 163,1: 18x15
#394 @ 903,746: 18x14
#395 @ 541,66: 19x16
#396 @ 340,167: 12x16
#397 @ 458,576: 23x23
#398 @ 54,511: 27x18
#399 @ 128,584: 25x12
#400 @ 71,283: 16x21
#401 @ 71,424: 23x12
#402 @ 438,61: 18x26
#403 @ 325,455: 19x24
#404 @ 176,73: 29x29
#405 @ 918,785: 17x23
#406 @ 36,232: 28x21
#407 @ 166,385: 25x25
#408 @ 125,953: 12x17
#409 @ 633,118: 29x26
#410 @ 32,83: 24x29
#411 @ 46,413: 19x14
#412 @ 416,763: 11x23
#413 @ 50,229: 26x11
#414 @ 414,97: 13x17
#415 @ 137,88: 10x14
#416 @ 914,701: 20x20
#417 @ 907,885: 17x27
#418 @ 665,202: 27x21
#419 @ 3,833: 11x25
#420 @ 185,932: 17x20
#421 @ 908,79: 29x19
#422 @ 389,83: 26x26
#423 @ 893,703: 14x29
#424 @ 915,881: 12x9
#425 @ 652,673: 12x11
#426 @ 726,221: 22x28
#427 @ 745,73: 28x15
#428 @ 682,856: 16x15
#429 @ 940,417: 29x17
#430 @ 594,796: 25x16
#431 @ 939,227: 17x18
#432 @ 707,983: 20x11
#433 @ 517,173: 20x21
#434 @ 719,392: 18x14
#435 @ 227,279: 11x20
#436 @ 4,655: 21x13
#437 @ 131,418: 28x24
#438 @ 123,724: 25x28
#439 @ 300,223: 18x28
#440 @ 184,807: 19x25
#441 @ 586,689: 17x12
#442 @ 583,400: 16x16
#443 @ 341,474: 25x23
#444 @ 725,346: 13x15
#445 @ 472,461: 18x10
#446 @ 679,136: 10x12
#447 @ 467,32: 28x21
#448 @ 256,348: 16x15
#449 @ 110,275: 28x15
#450 @ 254,812: 21x28
#451 @ 320,126: 27x25
#452 @ 510,295: 25x17
#453 @ 134,954: 22x28
#454 @ 333,137: 25x14
#455 @ 5,489: 20x16
#456 @ 523,51: 18x25
#457 @ 83,651: 21x24
#458 @ 97,652: 11x22
#459 @ 659,127: 29x23
#460 @ 569,336: 16x19
#461 @ 522,968: 21x23
#462 @ 700,412: 22x21
#463 @ 897,722: 25x29
#464 @ 193,577: 19x17
#465 @ 274,233: 14x12
#466 @ 201,140: 21x22
#467 @ 741,219: 17x12
#468 @ 515,388: 16x18
#469 @ 215,239: 29x10
#470 @ 435,724: 12x15
#471 @ 524,442: 10x29
#472 @ 11,84: 19x19
#473 @ 170,393: 29x29
#474 @ 278,404: 18x13
#475 @ 152,683: 18x21
#476 @ 608,637: 10x10
#477 @ 920,652: 27x29
#478 @ 301,218: 12x27
#479 @ 711,364: 10x18
#480 @ 919,804: 17x28
#481 @ 632,848: 16x15
#482 @ 972,802: 4x8
#483 @ 603,391: 25x10
#484 @ 577,263: 14x17
#485 @ 618,909: 21x10
#486 @ 905,91: 10x15
#487 @ 540,68: 13x29
#488 @ 51,660: 17x26
#489 @ 695,384: 11x16
#490 @ 208,360: 16x25
#491 @ 665,732: 15x18
#492 @ 447,841: 19x12
#493 @ 968,257: 14x20
#494 @ 651,815: 25x18
#495 @ 9,308: 10x17
#496 @ 320,538: 27x13
#497 @ 980,965: 18x20
#498 @ 402,132: 15x18
#499 @ 949,187: 24x12
#500 @ 403,561: 12x29
#501 @ 15,850: 8x12
#502 @ 50,619: 20x10
#503 @ 388,498: 27x28
#504 @ 114,209: 10x23
#505 @ 133,854: 21x19
#506 @ 33,809: 15x14
#507 @ 0,596: 13x18
#508 @ 120,533: 14x12
#509 @ 562,488: 13x13
#510 @ 292,771: 10x14
#511 @ 532,36: 24x16
#512 @ 202,207: 16x13
#513 @ 509,889: 25x26
#514 @ 703,965: 24x22
#515 @ 91,754: 11x15
#516 @ 208,354: 29x27
#517 @ 43,639: 10x14
#518 @ 232,184: 14x25
#519 @ 506,522: 16x11
#520 @ 366,282: 19x10
#521 @ 831,269: 18x15
#522 @ 406,665: 22x12
#523 @ 261,560: 11x15
#524 @ 585,306: 13x23
#525 @ 858,818: 27x12
#526 @ 919,370: 29x26
#527 @ 119,736: 21x27
#528 @ 931,880: 18x22
#529 @ 416,460: 18x14
#530 @ 132,258: 18x18
#531 @ 506,311: 17x11
#532 @ 128,344: 23x17
#533 @ 642,564: 23x13
#534 @ 170,226: 20x29
#535 @ 496,381: 23x11
#536 @ 187,809: 12x18
#537 @ 94,761: 17x10
#538 @ 949,242: 12x22
#539 @ 639,388: 28x22
#540 @ 492,268: 15x15
#541 @ 475,875: 24x27
#542 @ 851,817: 19x16
#543 @ 184,433: 23x12
#544 @ 690,388: 29x29
#545 @ 251,823: 17x27
#546 @ 726,130: 25x13
#547 @ 829,946: 22x16
#548 @ 875,66: 11x23
#549 @ 892,147: 17x22
#550 @ 40,100: 27x27
#551 @ 564,888: 13x14
#552 @ 620,546: 18x23
#553 @ 987,909: 11x13
#554 @ 253,937: 11x19
#555 @ 962,331: 28x21
#556 @ 387,701: 13x21
#557 @ 220,836: 21x25
#558 @ 34,794: 25x29
#559 @ 482,977: 20x16
#560 @ 441,848: 12x22
#561 @ 653,682: 21x18
#562 @ 139,214: 15x14
#563 @ 777,688: 14x21
#564 @ 219,551: 10x18
#565 @ 858,72: 12x11
#566 @ 495,580: 20x22
#567 @ 975,711: 24x21
#568 @ 747,793: 15x17
#569 @ 481,566: 24x23
#570 @ 145,700: 24x11
#571 @ 194,238: 10x17
#572 @ 396,509: 5x13
#573 @ 196,267: 10x26
#574 @ 238,556: 14x26
#575 @ 28,698: 25x23
#576 @ 110,732: 17x12
#577 @ 569,126: 29x16
#578 @ 126,727: 29x18
#579 @ 61,247: 27x18
#580 @ 91,603: 15x13
#581 @ 944,393: 18x13
#582 @ 824,729: 23x21
#583 @ 128,847: 28x13
#584 @ 865,72: 18x23
#585 @ 382,290: 16x10
#586 @ 874,148: 25x24
#587 @ 221,412: 25x26
#588 @ 769,379: 23x29
#589 @ 595,393: 16x26
#590 @ 52,246: 26x17
#591 @ 257,923: 11x18
#592 @ 956,966: 22x23
#593 @ 869,130: 13x25
#594 @ 554,495: 12x23
#595 @ 563,274: 21x25
#596 @ 256,571: 22x18
#597 @ 419,101: 14x17
#598 @ 192,599: 23x27
#599 @ 177,541: 15x25
#600 @ 48,293: 17x28
#601 @ 31,106: 23x13
#602 @ 678,875: 12x20
#603 @ 305,156: 11x18
#604 @ 172,125: 21x27
#605 @ 809,699: 25x15
#606 @ 882,67: 24x14
#607 @ 659,474: 21x13
#608 @ 688,954: 11x14
#609 @ 32,557: 28x27
#610 @ 38,800: 13x13
#611 @ 217,256: 14x26
#612 @ 605,367: 15x25
#613 @ 596,146: 22x14
#614 @ 30,812: 24x28
#615 @ 551,153: 23x16
#616 @ 790,784: 21x10
#617 @ 942,896: 23x11
#618 @ 355,844: 19x20
#619 @ 39,541: 22x22
#620 @ 434,497: 11x12
#621 @ 682,913: 26x22
#622 @ 323,589: 20x12
#623 @ 535,711: 19x21
#624 @ 750,207: 25x28
#625 @ 534,58: 20x15
#626 @ 323,18: 17x16
#627 @ 110,740: 25x15
#628 @ 868,211: 14x15
#629 @ 212,887: 24x28
#630 @ 438,733: 28x14
#631 @ 959,545: 13x21
#632 @ 755,887: 21x14
#633 @ 561,100: 19x11
#634 @ 192,651: 16x20
#635 @ 556,168: 27x14
#636 @ 905,883: 19x12
#637 @ 846,323: 15x10
#638 @ 886,296: 12x10
#639 @ 912,221: 24x29
#640 @ 695,373: 24x16
#641 @ 427,432: 27x11
#642 @ 956,797: 28x20
#643 @ 870,826: 13x13
#644 @ 891,373: 26x27
#645 @ 900,681: 21x19
#646 @ 132,239: 25x28
#647 @ 342,245: 15x9
#648 @ 257,544: 17x19
#649 @ 57,418: 15x16
#650 @ 338,329: 11x16
#651 @ 864,941: 15x16
#652 @ 280,720: 28x26
#653 @ 336,525: 22x17
#654 @ 870,496: 18x13
#655 @ 939,781: 23x22
#656 @ 925,445: 15x17
#657 @ 506,270: 13x13
#658 @ 962,641: 11x21
#659 @ 781,779: 16x14
#660 @ 693,158: 27x16
#661 @ 464,911: 19x26
#662 @ 534,452: 29x28
#663 @ 575,398: 13x19
#664 @ 315,843: 18x20
#665 @ 721,163: 25x13
#666 @ 784,526: 22x15
#667 @ 350,809: 19x25
#668 @ 351,43: 20x25
#669 @ 943,926: 24x22
#670 @ 511,308: 14x25
#671 @ 547,790: 29x22
#672 @ 801,41: 15x22
#673 @ 235,900: 13x12
#674 @ 871,486: 13x17
#675 @ 855,206: 20x18
#676 @ 762,670: 14x16
#677 @ 842,380: 16x18
#678 @ 140,741: 10x13
#679 @ 548,808: 28x10
#680 @ 595,758: 29x12
#681 @ 769,212: 18x24
#682 @ 884,337: 21x18
#683 @ 363,131: 22x10
#684 @ 622,285: 27x19
#685 @ 972,974: 21x13
#686 @ 881,885: 23x15
#687 @ 437,522: 15x21
#688 @ 235,359: 29x22
#689 @ 195,919: 22x25
#690 @ 84,594: 14x11
#691 @ 389,87: 26x18
#692 @ 887,909: 17x10
#693 @ 351,642: 13x19
#694 @ 10,474: 13x22
#695 @ 847,267: 16x27
#696 @ 728,436: 28x20
#697 @ 955,677: 26x14
#698 @ 24,796: 25x25
#699 @ 0,588: 22x28
#700 @ 654,306: 18x18
#701 @ 141,451: 16x19
#702 @ 575,336: 26x16
#703 @ 509,223: 20x25
#704 @ 104,328: 15x24
#705 @ 11,698: 23x16
#706 @ 985,851: 13x16
#707 @ 296,236: 20x24
#708 @ 669,206: 10x27
#709 @ 677,866: 17x22
#710 @ 336,530: 10x10
#711 @ 18,570: 29x13
#712 @ 272,524: 18x23
#713 @ 239,359: 20x20
#714 @ 583,180: 24x17
#715 @ 103,444: 16x13
#716 @ 946,0: 23x11
#717 @ 79,226: 22x24
#718 @ 159,801: 13x16
#719 @ 594,9: 20x22
#720 @ 453,387: 29x10
#721 @ 823,35: 22x13
#722 @ 616,16: 19x17
#723 @ 58,893: 26x11
#724 @ 853,811: 18x16
#725 @ 42,867: 25x22
#726 @ 758,139: 12x21
#727 @ 537,64: 27x14
#728 @ 121,208: 11x11
#729 @ 631,338: 12x23
#730 @ 691,808: 22x28
#731 @ 398,620: 21x15
#732 @ 47,206: 18x21
#733 @ 744,245: 23x21
#734 @ 797,16: 11x13
#735 @ 915,544: 24x14
#736 @ 845,724: 20x18
#737 @ 626,118: 13x24
#738 @ 566,929: 28x27
#739 @ 258,117: 23x24
#740 @ 680,338: 19x27
#741 @ 598,442: 23x22
#742 @ 340,530: 10x22
#743 @ 140,977: 28x16
#744 @ 317,824: 11x27
#745 @ 294,158: 19x11
#746 @ 130,973: 20x25
#747 @ 190,15: 11x15
#748 @ 191,198: 27x14
#749 @ 942,654: 20x24
#750 @ 510,803: 18x18
#751 @ 868,24: 26x15
#752 @ 827,856: 28x23
#753 @ 503,501: 27x26
#754 @ 875,467: 28x29
#755 @ 364,952: 29x25
#756 @ 948,156: 27x26
#757 @ 79,57: 10x29
#758 @ 380,97: 29x15
#759 @ 145,439: 20x23
#760 @ 683,343: 25x25
#761 @ 492,354: 18x27
#762 @ 812,434: 11x29
#763 @ 889,712: 22x23
#764 @ 216,804: 29x19
#765 @ 321,370: 23x24
#766 @ 637,188: 11x28
#767 @ 337,48: 23x26
#768 @ 296,219: 16x12
#769 @ 202,182: 29x26
#770 @ 888,350: 15x24
#771 @ 938,11: 21x17
#772 @ 348,206: 14x16
#773 @ 545,827: 23x23
#774 @ 954,888: 28x19
#775 @ 182,672: 22x13
#776 @ 884,871: 23x26
#777 @ 573,538: 23x28
#778 @ 42,98: 20x21
#779 @ 738,577: 19x24
#780 @ 836,966: 13x6
#781 @ 59,792: 18x11
#782 @ 160,626: 24x29
#783 @ 891,937: 17x17
#784 @ 773,65: 6x8
#785 @ 689,611: 17x22
#786 @ 770,153: 25x19
#787 @ 670,298: 17x27
#788 @ 924,891: 15x14
#789 @ 675,452: 12x13
#790 @ 761,186: 13x19
#791 @ 411,851: 13x11
#792 @ 959,450: 25x11
#793 @ 686,353: 10x21
#794 @ 945,559: 25x25
#795 @ 841,263: 24x14
#796 @ 892,490: 18x10
#797 @ 417,760: 10x28
#798 @ 107,321: 23x18
#799 @ 146,468: 17x17
#800 @ 428,422: 14x26
#801 @ 634,140: 28x14
#802 @ 194,141: 14x15
#803 @ 132,846: 13x15
#804 @ 261,372: 23x21
#805 @ 953,979: 25x16
#806 @ 914,880: 14x10
#807 @ 582,356: 11x10
#808 @ 50,626: 23x17
#809 @ 16,740: 23x22
#810 @ 91,754: 10x28
#811 @ 190,49: 26x18
#812 @ 808,809: 24x19
#813 @ 569,55: 13x28
#814 @ 773,791: 24x11
#815 @ 830,21: 13x28
#816 @ 902,495: 15x23
#817 @ 904,89: 26x21
#818 @ 220,350: 16x15
#819 @ 440,120: 23x13
#820 @ 69,216: 19x14
#821 @ 738,644: 27x20
#822 @ 98,77: 24x19
#823 @ 723,313: 16x24
#824 @ 510,786: 28x29
#825 @ 613,122: 25x21
#826 @ 794,56: 11x11
#827 @ 297,74: 18x13
#828 @ 417,380: 20x15
#829 @ 642,24: 18x13
#830 @ 815,683: 16x22
#831 @ 127,776: 25x17
#832 @ 373,249: 23x21
#833 @ 150,966: 13x24
#834 @ 239,417: 20x15
#835 @ 64,680: 12x12
#836 @ 197,200: 13x9
#837 @ 235,569: 24x19
#838 @ 69,693: 18x25
#839 @ 764,403: 26x29
#840 @ 163,836: 21x12
#841 @ 323,798: 13x11
#842 @ 504,371: 10x20
#843 @ 695,336: 17x26
#844 @ 554,82: 20x14
#845 @ 852,480: 29x14
#846 @ 418,90: 10x12
#847 @ 515,883: 23x24
#848 @ 983,681: 10x23
#849 @ 960,99: 25x17
#850 @ 482,905: 24x28
#851 @ 415,774: 12x28
#852 @ 668,792: 19x25
#853 @ 878,291: 19x28
#854 @ 53,590: 14x12
#855 @ 162,789: 25x13
#856 @ 859,535: 12x14
#857 @ 708,579: 12x10
#858 @ 748,477: 12x14
#859 @ 165,670: 23x18
#860 @ 297,985: 23x12
#861 @ 326,171: 24x10
#862 @ 610,610: 10x28
#863 @ 956,726: 22x25
#864 @ 27,809: 15x19
#865 @ 959,7: 10x12
#866 @ 838,923: 18x27
#867 @ 555,54: 5x7
#868 @ 243,421: 15x16
#869 @ 41,71: 11x27
#870 @ 637,835: 12x13
#871 @ 59,579: 24x16
#872 @ 916,875: 22x29
#873 @ 685,798: 24x21
#874 @ 892,58: 14x27
#875 @ 241,927: 16x21
#876 @ 425,630: 11x13
#877 @ 792,807: 21x29
#878 @ 311,583: 29x19
#879 @ 644,619: 15x18
#880 @ 112,579: 28x12
#881 @ 781,335: 25x19
#882 @ 586,742: 26x26
#883 @ 88,760: 25x22
#884 @ 675,434: 11x10
#885 @ 561,340: 21x11
#886 @ 172,658: 27x27
#887 @ 813,427: 14x10
#888 @ 696,577: 19x13
#889 @ 432,331: 20x16
#890 @ 158,796: 16x19
#891 @ 975,721: 12x20
#892 @ 127,87: 19x27
#893 @ 509,211: 14x26
#894 @ 627,419: 27x11
#895 @ 573,910: 28x23
#896 @ 84,662: 27x14
#897 @ 327,778: 23x29
#898 @ 114,352: 20x26
#899 @ 245,592: 28x19
#900 @ 45,291: 25x20
#901 @ 609,266: 14x12
#902 @ 765,324: 21x12
#903 @ 8,518: 29x27
#904 @ 494,287: 23x26
#905 @ 357,376: 11x11
#906 @ 626,395: 14x24
#907 @ 305,327: 11x28
#908 @ 420,846: 14x12
#909 @ 144,381: 11x26
#910 @ 443,378: 20x19
#911 @ 219,316: 9x5
#912 @ 961,738: 10x14
#913 @ 524,43: 27x20
#914 @ 246,597: 17x13
#915 @ 132,901: 19x24
#916 @ 266,123: 14x11
#917 @ 784,92: 12x11
#918 @ 140,23: 16x10
#919 @ 286,461: 14x21
#920 @ 378,558: 26x19
#921 @ 260,726: 28x11
#922 @ 736,446: 27x11
#923 @ 494,181: 11x29
#924 @ 658,835: 16x23
#925 @ 891,291: 23x16
#926 @ 218,356: 28x28
#927 @ 186,637: 28x18
#928 @ 96,26: 29x19
#929 @ 21,792: 29x12
#930 @ 894,734: 18x28
#931 @ 410,775: 24x11
#932 @ 39,369: 16x19
#933 @ 916,22: 23x17
#934 @ 661,99: 21x26
#935 @ 322,108: 24x22
#936 @ 715,157: 26x22
#937 @ 234,892: 14x21
#938 @ 489,453: 12x18
#939 @ 910,219: 26x17
#940 @ 285,756: 12x25
#941 @ 408,851: 20x15
#942 @ 584,44: 25x21
#943 @ 60,880: 20x16
#944 @ 66,727: 11x18
#945 @ 842,275: 21x18
#946 @ 983,282: 11x20
#947 @ 568,397: 22x15
#948 @ 966,748: 28x22
#949 @ 953,396: 19x12
#950 @ 348,791: 16x27
#951 @ 473,937: 12x20
#952 @ 237,652: 15x22
#953 @ 889,289: 20x10
#954 @ 391,93: 16x16
#955 @ 208,525: 29x22
#956 @ 675,365: 18x18
#957 @ 116,71: 19x21
#958 @ 43,522: 22x26
#959 @ 317,527: 22x15
#960 @ 237,204: 15x29
#961 @ 282,194: 17x15
#962 @ 598,299: 25x22
#963 @ 492,531: 21x16
#964 @ 249,637: 15x10
#965 @ 335,651: 19x26
#966 @ 735,839: 20x21
#967 @ 213,312: 26x14
#968 @ 205,209: 13x12
#969 @ 632,824: 15x18
#970 @ 213,204: 20x24
#971 @ 585,27: 15x27
#972 @ 469,933: 21x28
#973 @ 907,511: 29x12
#974 @ 534,725: 14x11
#975 @ 756,863: 22x29
#976 @ 715,2: 18x19
#977 @ 776,509: 27x18
#978 @ 665,433: 17x13
#979 @ 393,344: 16x14
#980 @ 819,968: 25x23
#981 @ 137,776: 10x16
#982 @ 128,821: 12x28
#983 @ 22,87: 12x14
#984 @ 357,203: 18x25
#985 @ 128,892: 24x22
#986 @ 988,833: 10x27
#987 @ 638,827: 24x21
#988 @ 860,952: 15x16
#989 @ 325,383: 24x25
#990 @ 593,221: 4x4
#991 @ 870,707: 20x10
#992 @ 905,364: 29x28
#993 @ 628,204: 27x29
#994 @ 678,206: 29x28
#995 @ 723,281: 23x27
#996 @ 970,184: 13x16
#997 @ 175,435: 19x13
#998 @ 785,395: 23x21
#999 @ 341,116: 22x24
#1000 @ 140,496: 24x26
#1001 @ 70,500: 18x24
#1002 @ 637,363: 10x13
#1003 @ 431,769: 12x10
#1004 @ 621,609: 29x18
#1005 @ 335,526: 21x28
#1006 @ 692,862: 18x14
#1007 @ 720,706: 17x20
#1008 @ 765,145: 28x11
#1009 @ 409,818: 17x12
#1010 @ 83,515: 18x10
#1011 @ 215,278: 11x22
#1012 @ 786,78: 10x28
#1013 @ 608,43: 23x17
#1014 @ 158,758: 15x21
#1015 @ 77,659: 13x27
#1016 @ 884,73: 27x19
#1017 @ 181,566: 13x19
#1018 @ 302,745: 21x24
#1019 @ 198,537: 11x19
#1020 @ 594,361: 21x15
#1021 @ 532,642: 28x14
#1022 @ 690,488: 24x20
#1023 @ 517,900: 17x28
#1024 @ 8,848: 19x17
#1025 @ 834,963: 23x19
#1026 @ 572,188: 15x12
#1027 @ 439,489: 20x15
#1028 @ 773,271: 28x11
#1029 @ 19,360: 20x24
#1030 @ 111,128: 23x13
#1031 @ 203,314: 25x27
#1032 @ 175,693: 16x12
#1033 @ 100,636: 12x25
#1034 @ 449,40: 28x27
#1035 @ 894,65: 13x22
#1036 @ 527,174: 14x24
#1037 @ 745,186: 27x15
#1038 @ 838,720: 18x11
#1039 @ 96,764: 20x16
#1040 @ 158,966: 25x26
#1041 @ 816,158: 14x27
#1042 @ 879,100: 20x16
#1043 @ 121,332: 17x29
#1044 @ 681,626: 16x27
#1045 @ 853,122: 20x26
#1046 @ 642,841: 19x10
#1047 @ 224,68: 27x19
#1048 @ 886,950: 28x21
#1049 @ 892,65: 22x26
#1050 @ 125,16: 16x11
#1051 @ 250,725: 21x20
#1052 @ 284,393: 10x14
#1053 @ 777,408: 15x21
#1054 @ 968,633: 20x18
#1055 @ 806,411: 10x13
#1056 @ 431,498: 22x22
#1057 @ 913,267: 16x27
#1058 @ 432,715: 11x12
#1059 @ 413,868: 24x24
#1060 @ 360,338: 19x27
#1061 @ 539,136: 28x20
#1062 @ 639,591: 20x29
#1063 @ 250,402: 19x21
#1064 @ 339,16: 14x23
#1065 @ 97,602: 13x10
#1066 @ 850,818: 10x14
#1067 @ 221,231: 20x29
#1068 @ 780,79: 29x25
#1069 @ 825,711: 28x23
#1070 @ 452,724: 16x14
#1071 @ 599,441: 11x17
#1072 @ 302,773: 17x19
#1073 @ 954,455: 10x26
#1074 @ 42,646: 18x20
#1075 @ 851,362: 27x19
#1076 @ 40,513: 10x18
#1077 @ 273,805: 12x25
#1078 @ 884,688: 23x14
#1079 @ 407,848: 19x18
#1080 @ 590,541: 13x28
#1081 @ 417,582: 17x18
#1082 @ 244,931: 15x15
#1083 @ 753,566: 12x27
#1084 @ 47,89: 19x14
#1085 @ 702,745: 27x12
#1086 @ 22,704: 26x17
#1087 @ 740,305: 16x28
#1088 @ 395,853: 21x17
#1089 @ 926,408: 16x29
#1090 @ 250,814: 23x24
#1091 @ 636,344: 16x15
#1092 @ 882,671: 12x17
#1093 @ 390,460: 23x16
#1094 @ 216,149: 19x11
#1095 @ 115,513: 17x27
#1096 @ 142,621: 16x20
#1097 @ 218,984: 14x15
#1098 @ 641,613: 25x29
#1099 @ 466,73: 19x25
#1100 @ 686,644: 16x21
#1101 @ 237,164: 23x11
#1102 @ 285,161: 10x29
#1103 @ 810,654: 28x13
#1104 @ 774,93: 11x13
#1105 @ 517,425: 15x25
#1106 @ 85,143: 14x10
#1107 @ 521,430: 17x18
#1108 @ 150,223: 19x27
#1109 @ 966,238: 20x26
#1110 @ 599,403: 29x22
#1111 @ 472,269: 19x25
#1112 @ 8,33: 24x23
#1113 @ 944,907: 17x27
#1114 @ 935,685: 11x13
#1115 @ 224,507: 29x21
#1116 @ 692,457: 24x23
#1117 @ 26,776: 16x26
#1118 @ 272,147: 28x27
#1119 @ 449,531: 15x23
#1120 @ 912,871: 23x10
#1121 @ 640,619: 22x12
#1122 @ 1,727: 15x24
#1123 @ 407,338: 12x17
#1124 @ 140,493: 27x17
#1125 @ 465,938: 10x11
#1126 @ 129,489: 25x25
#1127 @ 482,969: 11x28
#1128 @ 793,10: 25x26
#1129 @ 489,376: 19x17
#1130 @ 440,917: 28x22
#1131 @ 572,926: 22x11
#1132 @ 234,564: 29x10
#1133 @ 29,211: 26x21
#1134 @ 641,851: 21x26
#1135 @ 757,77: 22x23
#1136 @ 683,272: 26x17
#1137 @ 508,367: 15x14
#1138 @ 684,885: 12x16
#1139 @ 935,900: 9x9
#1140 @ 438,507: 11x16
#1141 @ 143,416: 11x11
#1142 @ 217,280: 6x13
#1143 @ 183,240: 27x17
#1144 @ 439,777: 22x19
#1145 @ 534,221: 27x13
#1146 @ 834,971: 16x19
#1147 @ 874,900: 19x15
#1148 @ 719,392: 14x25
#1149 @ 903,925: 24x18
#1150 @ 2,361: 20x14
#1151 @ 325,841: 17x18
#1152 @ 242,718: 13x20
#1153 @ 647,835: 14x27
#1154 @ 72,408: 23x28
#1155 @ 624,546: 29x21
#1156 @ 925,624: 15x16
#1157 @ 445,45: 12x23
#1158 @ 36,867: 18x20
#1159 @ 954,691: 16x25
#1160 @ 383,106: 14x18
#1161 @ 699,474: 14x25
#1162 @ 837,324: 18x14
#1163 @ 677,709: 14x19
#1164 @ 220,231: 13x16
#1165 @ 633,895: 21x16
#1166 @ 615,565: 22x12
#1167 @ 389,713: 8x3
#1168 @ 500,899: 23x11
#1169 @ 252,429: 27x21
#1170 @ 262,357: 24x19
#1171 @ 948,10: 20x17
#1172 @ 981,317: 13x18
#1173 @ 940,181: 23x17
#1174 @ 397,917: 10x29
#1175 @ 947,704: 16x15
#1176 @ 14,227: 25x10
#1177 @ 130,257: 19x17
#1178 @ 973,909: 20x26
#1179 @ 593,939: 24x11
#1180 @ 931,897: 17x16
#1181 @ 669,847: 20x26
#1182 @ 500,60: 12x23
#1183 @ 975,97: 15x13
#1184 @ 124,462: 21x12
#1185 @ 469,858: 10x22
#1186 @ 84,75: 12x14
#1187 @ 402,853: 16x18
#1188 @ 860,79: 24x28
#1189 @ 133,404: 23x21
#1190 @ 533,717: 16x26
#1191 @ 383,610: 26x13
#1192 @ 594,261: 16x21
#1193 @ 501,433: 29x11
#1194 @ 240,560: 27x13
#1195 @ 364,550: 19x21
#1196 @ 907,858: 26x25
#1197 @ 680,852: 13x19
#1198 @ 829,172: 12x26
#1199 @ 961,178: 16x23
#1200 @ 458,283: 23x23
#1201 @ 804,904: 12x25
#1202 @ 148,471: 14x28
#1203 @ 207,722: 10x24
#1204 @ 388,877: 28x29
#1205 @ 895,932: 23x21
#1206 @ 308,203: 28x10
#1207 @ 460,550: 19x16
#1208 @ 153,752: 26x12
#1209 @ 632,290: 20x13
#1210 @ 223,552: 28x13
#1211 @ 589,151: 17x22
#1212 @ 850,810: 17x24
#1213 @ 331,630: 17x6
#1214 @ 528,51: 29x22
#1215 @ 91,32: 29x15
#1216 @ 193,550: 14x29
#1217 @ 771,63: 12x17
#1218 @ 74,589: 26x11
#1219 @ 840,262: 25x21
#1220 @ 215,768: 17x11
#1221 @ 904,933: 20x20
#1222 @ 188,141: 11x20
#1223 @ 509,397: 18x27
#1224 @ 220,546: 15x12
#1225 @ 962,536: 14x21
#1226 @ 911,390: 26x24
#1227 @ 610,132: 25x25
#1228 @ 444,846: 20x21
#1229 @ 912,803: 13x21
#1230 @ 165,806: 12x27
#1231 @ 373,248: 17x23
#1232 @ 807,896: 13x19
#1233 @ 520,554: 12x27
#1234 @ 94,511: 19x11
#1235 @ 435,689: 17x11
#1236 @ 665,653: 13x13
#1237 @ 130,242: 11x19
#1238 @ 530,216: 21x15
#1239 @ 635,575: 27x10
#1240 @ 12,692: 25x24
#1241 @ 137,234: 25x25
#1242 @ 438,628: 21x20
#1243 @ 257,935: 23x10
#1244 @ 275,246: 14x22
#1245 @ 383,80: 20x27
#1246 @ 194,82: 27x15
#1247 @ 836,864: 18x12
#1248 @ 797,687: 21x11
#1249 @ 4,114: 25x21
#1250 @ 725,543: 13x23
#1251 @ 710,524: 10x17
#1252 @ 49,472: 14x18
#1253 @ 928,883: 17x19
#1254 @ 435,127: 26x18
#1255 @ 913,708: 17x18
#1256 @ 248,564: 19x14
#1257 @ 827,713: 22x18
#1258 @ 883,72: 13x16
#1259 @ 220,318: 19x13
#1260 @ 416,126: 25x14
#1261 @ 130,103: 28x10
#1262 @ 712,824: 17x10
#1263 @ 448,49: 23x29
#1264 @ 344,403: 18x23
#1265 @ 277,256: 3x8
#1266 @ 81,179: 15x25
#1267 @ 206,519: 22x23
#1268 @ 902,277: 15x23
#1269 @ 673,443: 18x26
#1270 @ 726,331: 22x18
#1271 @ 73,170: 19x14
#1272 @ 499,548: 12x14
#1273 @ 587,802: 29x22
#1274 @ 77,644: 24x27
#1275 @ 469,828: 14x3
#1276 @ 888,483: 26x10
#1277 @ 50,360: 17x27
#1278 @ 857,790: 21x29
#1279 @ 231,894: 16x12
#1280 @ 23,131: 27x29
#1281 @ 73,300: 10x12
#1282 @ 286,137: 20x25
#1283 @ 673,866: 21x27
#1284 @ 347,704: 27x23
#1285 @ 445,504: 16x25
#1286 @ 919,668: 28x20
#1287 @ 74,729: 29x11
#1288 @ 957,313: 21x13
#1289 @ 566,873: 18x16
#1290 @ 407,547: 28x17
#1291 @ 581,354: 23x22
#1292 @ 119,762: 20x28
#1293 @ 174,133: 29x26
#1294 @ 265,262: 14x19
#1295 @ 548,129: 28x11
#1296 @ 57,32: 10x29
#1297 @ 780,123: 14x28
#1298 @ 195,339: 17x20
#1299 @ 892,193: 28x11
#1300 @ 634,811: 17x27
#1301 @ 130,498: 26x10`

	claims := strings.Split(claims_raw, "\n")

	var claimTs []claimT

	cells := make(map[string]int)

	for _, claim := range claims {
		entry := strings.Split(claim, "@")[1]
		entry_parts := strings.Split(entry, ":")
		coords := entry_parts[0]
		coords_parts := strings.Split(coords, ",")
		col, _ := strconv.Atoi(strings.TrimSpace(coords_parts[0]))
		row, _ := strconv.Atoi(strings.TrimSpace(coords_parts[1]))
		dimens := entry_parts[1]
		dimens_parts := strings.Split(dimens, "x")
		width, _ := strconv.Atoi(strings.TrimSpace(dimens_parts[0]))
		height, _ := strconv.Atoi(strings.TrimSpace(dimens_parts[1]))
		newClaim := claimT{row: row, col: col, width: width, height: height}
		claimTs = append(claimTs, newClaim)
	}

	for _, cl := range claimTs {
		for r := cl.row ; r < cl.row + cl.height ; r++ {
			for c := cl.col ; c < cl.col + cl.width ; c++ {
				key := fmt.Sprintf("%d_%d", r, c)
				cells[key]++
			}
		}
	}

	two_claims := 0
	for _, value := range cells {
		if value >= 2 {
			two_claims++
		}
	}

	fmt.Println(two_claims)

	for id, cl := range claimTs {
		only_one := true
		for r := cl.row ; r < cl.row + cl.height ; r++ {
			for c := cl.col ; c < cl.col + cl.width ; c++ {
				key := fmt.Sprintf("%d_%d", r, c)
				value, _ := cells[key]
				if value > 1 {
					only_one = false
					break
				}
			}
			if !only_one {
				break
			}
		}
		if only_one {
			fmt.Println(id+1)
			//break
		}
	}
}
