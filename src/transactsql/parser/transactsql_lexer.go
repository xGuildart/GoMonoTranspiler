// Code generated from TransactSQL.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 68, 711,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 7, 11, 213, 10, 11, 12, 11, 14, 11, 216, 11, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3,
	48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3,
	54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55,
	3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3,
	56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58,
	3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3,
	59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 7, 61,
	639, 10, 61, 12, 61, 14, 61, 642, 11, 61, 3, 61, 5, 61, 645, 10, 61, 3,
	61, 3, 61, 3, 62, 3, 62, 7, 62, 651, 10, 62, 12, 62, 14, 62, 654, 11, 62,
	3, 63, 3, 63, 5, 63, 658, 10, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3,
	66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71,
	3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 73, 6, 73, 683, 10, 73, 13,
	73, 14, 73, 684, 3, 73, 3, 73, 3, 74, 3, 74, 5, 74, 691, 10, 74, 3, 74,
	5, 74, 694, 10, 74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 75, 3, 75, 7, 75, 702,
	10, 75, 12, 75, 14, 75, 705, 11, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75,
	3, 703, 2, 76, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73,
	38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91,
	47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55,
	109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 2,
	125, 63, 127, 2, 129, 2, 131, 2, 133, 2, 135, 2, 137, 2, 139, 2, 141, 64,
	143, 65, 145, 66, 147, 67, 149, 68, 3, 2, 10, 4, 2, 67, 92, 99, 124, 5,
	2, 50, 59, 67, 92, 99, 124, 3, 2, 97, 97, 3, 2, 51, 59, 3, 2, 50, 59, 12,
	2, 36, 36, 41, 41, 65, 65, 94, 94, 99, 100, 104, 104, 112, 112, 116, 116,
	118, 118, 120, 120, 3, 2, 34, 34, 4, 2, 11, 11, 34, 34, 2, 711, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3,
	2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65,
	3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2,
	73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2,
	2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2,
	2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2,
	2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103,
	3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2,
	2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3,
	2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2,
	141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2,
	2, 2, 2, 149, 3, 2, 2, 2, 3, 151, 3, 2, 2, 2, 5, 158, 3, 2, 2, 2, 7, 160,
	3, 2, 2, 2, 9, 170, 3, 2, 2, 2, 11, 179, 3, 2, 2, 2, 13, 181, 3, 2, 2,
	2, 15, 186, 3, 2, 2, 2, 17, 188, 3, 2, 2, 2, 19, 206, 3, 2, 2, 2, 21, 210,
	3, 2, 2, 2, 23, 217, 3, 2, 2, 2, 25, 221, 3, 2, 2, 2, 27, 238, 3, 2, 2,
	2, 29, 254, 3, 2, 2, 2, 31, 260, 3, 2, 2, 2, 33, 271, 3, 2, 2, 2, 35, 275,
	3, 2, 2, 2, 37, 281, 3, 2, 2, 2, 39, 288, 3, 2, 2, 2, 41, 298, 3, 2, 2,
	2, 43, 309, 3, 2, 2, 2, 45, 316, 3, 2, 2, 2, 47, 324, 3, 2, 2, 2, 49, 341,
	3, 2, 2, 2, 51, 356, 3, 2, 2, 2, 53, 361, 3, 2, 2, 2, 55, 363, 3, 2, 2,
	2, 57, 367, 3, 2, 2, 2, 59, 375, 3, 2, 2, 2, 61, 387, 3, 2, 2, 2, 63, 402,
	3, 2, 2, 2, 65, 408, 3, 2, 2, 2, 67, 412, 3, 2, 2, 2, 69, 414, 3, 2, 2,
	2, 71, 416, 3, 2, 2, 2, 73, 422, 3, 2, 2, 2, 75, 430, 3, 2, 2, 2, 77, 438,
	3, 2, 2, 2, 79, 443, 3, 2, 2, 2, 81, 447, 3, 2, 2, 2, 83, 460, 3, 2, 2,
	2, 85, 463, 3, 2, 2, 2, 87, 491, 3, 2, 2, 2, 89, 501, 3, 2, 2, 2, 91, 509,
	3, 2, 2, 2, 93, 520, 3, 2, 2, 2, 95, 522, 3, 2, 2, 2, 97, 524, 3, 2, 2,
	2, 99, 531, 3, 2, 2, 2, 101, 535, 3, 2, 2, 2, 103, 539, 3, 2, 2, 2, 105,
	562, 3, 2, 2, 2, 107, 585, 3, 2, 2, 2, 109, 591, 3, 2, 2, 2, 111, 604,
	3, 2, 2, 2, 113, 611, 3, 2, 2, 2, 115, 615, 3, 2, 2, 2, 117, 620, 3, 2,
	2, 2, 119, 630, 3, 2, 2, 2, 121, 635, 3, 2, 2, 2, 123, 648, 3, 2, 2, 2,
	125, 657, 3, 2, 2, 2, 127, 659, 3, 2, 2, 2, 129, 661, 3, 2, 2, 2, 131,
	663, 3, 2, 2, 2, 133, 665, 3, 2, 2, 2, 135, 667, 3, 2, 2, 2, 137, 669,
	3, 2, 2, 2, 139, 672, 3, 2, 2, 2, 141, 674, 3, 2, 2, 2, 143, 677, 3, 2,
	2, 2, 145, 682, 3, 2, 2, 2, 147, 693, 3, 2, 2, 2, 149, 697, 3, 2, 2, 2,
	151, 152, 7, 85, 2, 2, 152, 153, 7, 69, 2, 2, 153, 154, 7, 74, 2, 2, 154,
	155, 7, 71, 2, 2, 155, 156, 7, 79, 2, 2, 156, 157, 7, 67, 2, 2, 157, 4,
	3, 2, 2, 2, 158, 159, 7, 46, 2, 2, 159, 6, 3, 2, 2, 2, 160, 161, 7, 93,
	2, 2, 161, 162, 7, 82, 2, 2, 162, 163, 7, 84, 2, 2, 163, 164, 7, 75, 2,
	2, 164, 165, 7, 79, 2, 2, 165, 166, 7, 67, 2, 2, 166, 167, 7, 84, 2, 2,
	167, 168, 7, 91, 2, 2, 168, 169, 7, 95, 2, 2, 169, 8, 3, 2, 2, 2, 170,
	171, 7, 36, 2, 2, 171, 172, 7, 102, 2, 2, 172, 173, 7, 103, 2, 2, 173,
	174, 7, 104, 2, 2, 174, 175, 7, 99, 2, 2, 175, 176, 7, 119, 2, 2, 176,
	177, 7, 110, 2, 2, 177, 178, 7, 36, 2, 2, 178, 10, 3, 2, 2, 2, 179, 180,
	7, 48, 2, 2, 180, 12, 3, 2, 2, 2, 181, 182, 7, 80, 2, 2, 182, 183, 7, 81,
	2, 2, 183, 184, 7, 80, 2, 2, 184, 185, 7, 71, 2, 2, 185, 14, 3, 2, 2, 2,
	186, 187, 7, 61, 2, 2, 187, 16, 3, 2, 2, 2, 188, 189, 7, 83, 2, 2, 189,
	190, 7, 87, 2, 2, 190, 191, 7, 81, 2, 2, 191, 192, 7, 86, 2, 2, 192, 193,
	7, 71, 2, 2, 193, 194, 7, 70, 2, 2, 194, 195, 7, 97, 2, 2, 195, 196, 7,
	75, 2, 2, 196, 197, 7, 70, 2, 2, 197, 198, 7, 71, 2, 2, 198, 199, 7, 80,
	2, 2, 199, 200, 7, 86, 2, 2, 200, 201, 7, 75, 2, 2, 201, 202, 7, 72, 2,
	2, 202, 203, 7, 75, 2, 2, 203, 204, 7, 71, 2, 2, 204, 205, 7, 84, 2, 2,
	205, 18, 3, 2, 2, 2, 206, 207, 7, 81, 2, 2, 207, 208, 7, 72, 2, 2, 208,
	209, 7, 72, 2, 2, 209, 20, 3, 2, 2, 2, 210, 214, 5, 133, 67, 2, 211, 213,
	5, 135, 68, 2, 212, 211, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3,
	2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 22, 3, 2, 2, 2, 216, 214, 3, 2, 2,
	2, 217, 218, 7, 67, 2, 2, 218, 219, 7, 70, 2, 2, 219, 220, 7, 70, 2, 2,
	220, 24, 3, 2, 2, 2, 221, 222, 7, 67, 2, 2, 222, 223, 7, 78, 2, 2, 223,
	224, 7, 78, 2, 2, 224, 225, 7, 81, 2, 2, 225, 226, 7, 89, 2, 2, 226, 227,
	7, 97, 2, 2, 227, 228, 7, 82, 2, 2, 228, 229, 7, 67, 2, 2, 229, 230, 7,
	73, 2, 2, 230, 231, 7, 71, 2, 2, 231, 232, 7, 97, 2, 2, 232, 233, 7, 78,
	2, 2, 233, 234, 7, 81, 2, 2, 234, 235, 7, 69, 2, 2, 235, 236, 7, 77, 2,
	2, 236, 237, 7, 85, 2, 2, 237, 26, 3, 2, 2, 2, 238, 239, 7, 67, 2, 2, 239,
	240, 7, 78, 2, 2, 240, 241, 7, 78, 2, 2, 241, 242, 7, 81, 2, 2, 242, 243,
	7, 89, 2, 2, 243, 244, 7, 97, 2, 2, 244, 245, 7, 84, 2, 2, 245, 246, 7,
	81, 2, 2, 246, 247, 7, 89, 2, 2, 247, 248, 7, 97, 2, 2, 248, 249, 7, 78,
	2, 2, 249, 250, 7, 81, 2, 2, 250, 251, 7, 69, 2, 2, 251, 252, 7, 77, 2,
	2, 252, 253, 7, 85, 2, 2, 253, 28, 3, 2, 2, 2, 254, 255, 7, 67, 2, 2, 255,
	256, 7, 78, 2, 2, 256, 257, 7, 86, 2, 2, 257, 258, 7, 71, 2, 2, 258, 259,
	7, 84, 2, 2, 259, 30, 3, 2, 2, 2, 260, 261, 7, 67, 2, 2, 261, 262, 7, 80,
	2, 2, 262, 263, 7, 85, 2, 2, 263, 264, 7, 75, 2, 2, 264, 265, 7, 97, 2,
	2, 265, 266, 7, 80, 2, 2, 266, 267, 7, 87, 2, 2, 267, 268, 7, 78, 2, 2,
	268, 269, 7, 78, 2, 2, 269, 270, 7, 85, 2, 2, 270, 32, 3, 2, 2, 2, 271,
	272, 7, 67, 2, 2, 272, 273, 7, 85, 2, 2, 273, 274, 7, 69, 2, 2, 274, 34,
	3, 2, 2, 2, 275, 276, 7, 69, 2, 2, 276, 277, 7, 74, 2, 2, 277, 278, 7,
	71, 2, 2, 278, 279, 7, 69, 2, 2, 279, 280, 7, 77, 2, 2, 280, 36, 3, 2,
	2, 2, 281, 282, 7, 93, 2, 2, 282, 283, 7, 101, 2, 2, 283, 284, 7, 106,
	2, 2, 284, 285, 7, 99, 2, 2, 285, 286, 7, 116, 2, 2, 286, 287, 7, 95, 2,
	2, 287, 38, 3, 2, 2, 2, 288, 289, 7, 69, 2, 2, 289, 290, 7, 78, 2, 2, 290,
	291, 7, 87, 2, 2, 291, 292, 7, 85, 2, 2, 292, 293, 7, 86, 2, 2, 293, 294,
	7, 71, 2, 2, 294, 295, 7, 84, 2, 2, 295, 296, 7, 71, 2, 2, 296, 297, 7,
	70, 2, 2, 297, 40, 3, 2, 2, 2, 298, 299, 7, 69, 2, 2, 299, 300, 7, 81,
	2, 2, 300, 301, 7, 80, 2, 2, 301, 302, 7, 85, 2, 2, 302, 303, 7, 86, 2,
	2, 303, 304, 7, 84, 2, 2, 304, 305, 7, 67, 2, 2, 305, 306, 7, 75, 2, 2,
	306, 307, 7, 80, 2, 2, 307, 308, 7, 86, 2, 2, 308, 42, 3, 2, 2, 2, 309,
	310, 7, 69, 2, 2, 310, 311, 7, 84, 2, 2, 311, 312, 7, 71, 2, 2, 312, 313,
	7, 67, 2, 2, 313, 314, 7, 86, 2, 2, 314, 315, 7, 71, 2, 2, 315, 44, 3,
	2, 2, 2, 316, 317, 7, 102, 2, 2, 317, 318, 7, 103, 2, 2, 318, 319, 7, 104,
	2, 2, 319, 320, 7, 99, 2, 2, 320, 321, 7, 119, 2, 2, 321, 322, 7, 110,
	2, 2, 322, 323, 7, 118, 2, 2, 323, 46, 3, 2, 2, 2, 324, 325, 7, 70, 2,
	2, 325, 326, 7, 71, 2, 2, 326, 327, 7, 72, 2, 2, 327, 328, 7, 67, 2, 2,
	328, 329, 7, 87, 2, 2, 329, 330, 7, 78, 2, 2, 330, 331, 7, 86, 2, 2, 331,
	332, 7, 97, 2, 2, 332, 333, 7, 78, 2, 2, 333, 334, 7, 67, 2, 2, 334, 335,
	7, 80, 2, 2, 335, 336, 7, 73, 2, 2, 336, 337, 7, 87, 2, 2, 337, 338, 7,
	67, 2, 2, 338, 339, 7, 73, 2, 2, 339, 340, 7, 71, 2, 2, 340, 48, 3, 2,
	2, 2, 341, 342, 7, 70, 2, 2, 342, 343, 7, 71, 2, 2, 343, 344, 7, 72, 2,
	2, 344, 345, 7, 67, 2, 2, 345, 346, 7, 87, 2, 2, 346, 347, 7, 78, 2, 2,
	347, 348, 7, 86, 2, 2, 348, 349, 7, 97, 2, 2, 349, 350, 7, 85, 2, 2, 350,
	351, 7, 69, 2, 2, 351, 352, 7, 74, 2, 2, 352, 353, 7, 71, 2, 2, 353, 354,
	7, 79, 2, 2, 354, 355, 7, 67, 2, 2, 355, 50, 3, 2, 2, 2, 356, 357, 7, 70,
	2, 2, 357, 358, 7, 71, 2, 2, 358, 359, 7, 85, 2, 2, 359, 360, 7, 69, 2,
	2, 360, 52, 3, 2, 2, 2, 361, 362, 7, 63, 2, 2, 362, 54, 3, 2, 2, 2, 363,
	364, 7, 72, 2, 2, 364, 365, 7, 81, 2, 2, 365, 366, 7, 84, 2, 2, 366, 56,
	3, 2, 2, 2, 367, 368, 7, 72, 2, 2, 368, 369, 7, 81, 2, 2, 369, 370, 7,
	84, 2, 2, 370, 371, 7, 71, 2, 2, 371, 372, 7, 75, 2, 2, 372, 373, 7, 73,
	2, 2, 373, 374, 7, 80, 2, 2, 374, 58, 3, 2, 2, 2, 375, 376, 7, 93, 2, 2,
	376, 377, 7, 105, 2, 2, 377, 378, 7, 103, 2, 2, 378, 379, 7, 113, 2, 2,
	379, 380, 7, 105, 2, 2, 380, 381, 7, 116, 2, 2, 381, 382, 7, 99, 2, 2,
	382, 383, 7, 114, 2, 2, 383, 384, 7, 106, 2, 2, 384, 385, 7, 123, 2, 2,
	385, 386, 7, 95, 2, 2, 386, 60, 3, 2, 2, 2, 387, 388, 7, 75, 2, 2, 388,
	389, 7, 73, 2, 2, 389, 390, 7, 80, 2, 2, 390, 391, 7, 81, 2, 2, 391, 392,
	7, 84, 2, 2, 392, 393, 7, 71, 2, 2, 393, 394, 7, 97, 2, 2, 394, 395, 7,
	70, 2, 2, 395, 396, 7, 87, 2, 2, 396, 397, 7, 82, 2, 2, 397, 398, 7, 97,
	2, 2, 398, 399, 7, 77, 2, 2, 399, 400, 7, 71, 2, 2, 400, 401, 7, 91, 2,
	2, 401, 62, 3, 2, 2, 2, 402, 403, 7, 93, 2, 2, 403, 404, 7, 107, 2, 2,
	404, 405, 7, 112, 2, 2, 405, 406, 7, 118, 2, 2, 406, 407, 7, 95, 2, 2,
	407, 64, 3, 2, 2, 2, 408, 409, 7, 77, 2, 2, 409, 410, 7, 71, 2, 2, 410,
	411, 7, 91, 2, 2, 411, 66, 3, 2, 2, 2, 412, 413, 7, 42, 2, 2, 413, 68,
	3, 2, 2, 2, 414, 415, 7, 93, 2, 2, 415, 70, 3, 2, 2, 2, 416, 417, 7, 78,
	2, 2, 417, 418, 7, 81, 2, 2, 418, 419, 7, 73, 2, 2, 419, 420, 7, 75, 2,
	2, 420, 421, 7, 80, 2, 2, 421, 72, 3, 2, 2, 2, 422, 423, 7, 93, 2, 2, 423,
	424, 7, 111, 2, 2, 424, 425, 7, 113, 2, 2, 425, 426, 7, 112, 2, 2, 426,
	427, 7, 103, 2, 2, 427, 428, 7, 123, 2, 2, 428, 429, 7, 95, 2, 2, 429,
	74, 3, 2, 2, 2, 430, 431, 7, 80, 2, 2, 431, 432, 7, 81, 2, 2, 432, 433,
	7, 69, 2, 2, 433, 434, 7, 74, 2, 2, 434, 435, 7, 71, 2, 2, 435, 436, 7,
	69, 2, 2, 436, 437, 7, 77, 2, 2, 437, 76, 3, 2, 2, 2, 438, 439, 7, 80,
	2, 2, 439, 440, 7, 87, 2, 2, 440, 441, 7, 78, 2, 2, 441, 442, 7, 78, 2,
	2, 442, 78, 3, 2, 2, 2, 443, 444, 7, 80, 2, 2, 444, 445, 7, 81, 2, 2, 445,
	446, 7, 86, 2, 2, 446, 80, 3, 2, 2, 2, 447, 448, 7, 80, 2, 2, 448, 449,
	7, 81, 2, 2, 449, 450, 7, 80, 2, 2, 450, 451, 7, 69, 2, 2, 451, 452, 7,
	78, 2, 2, 452, 453, 7, 87, 2, 2, 453, 454, 7, 85, 2, 2, 454, 455, 7, 86,
	2, 2, 455, 456, 7, 71, 2, 2, 456, 457, 7, 84, 2, 2, 457, 458, 7, 71, 2,
	2, 458, 459, 7, 70, 2, 2, 459, 82, 3, 2, 2, 2, 460, 461, 7, 81, 2, 2, 461,
	462, 7, 80, 2, 2, 462, 84, 3, 2, 2, 2, 463, 464, 7, 81, 2, 2, 464, 465,
	7, 82, 2, 2, 465, 466, 7, 86, 2, 2, 466, 467, 7, 75, 2, 2, 467, 468, 7,
	79, 2, 2, 468, 469, 7, 75, 2, 2, 469, 470, 7, 92, 2, 2, 470, 471, 7, 71,
	2, 2, 471, 472, 7, 97, 2, 2, 472, 473, 7, 72, 2, 2, 473, 474, 7, 81, 2,
	2, 474, 475, 7, 84, 2, 2, 475, 476, 7, 97, 2, 2, 476, 477, 7, 85, 2, 2,
	477, 478, 7, 71, 2, 2, 478, 479, 7, 83, 2, 2, 479, 480, 7, 87, 2, 2, 480,
	481, 7, 71, 2, 2, 481, 482, 7, 80, 2, 2, 482, 483, 7, 86, 2, 2, 483, 484,
	7, 75, 2, 2, 484, 485, 7, 67, 2, 2, 485, 486, 7, 78, 2, 2, 486, 487, 7,
	97, 2, 2, 487, 488, 7, 77, 2, 2, 488, 489, 7, 71, 2, 2, 489, 490, 7, 91,
	2, 2, 490, 86, 3, 2, 2, 2, 491, 492, 7, 82, 2, 2, 492, 493, 7, 67, 2, 2,
	493, 494, 7, 70, 2, 2, 494, 495, 7, 97, 2, 2, 495, 496, 7, 75, 2, 2, 496,
	497, 7, 80, 2, 2, 497, 498, 7, 70, 2, 2, 498, 499, 7, 71, 2, 2, 499, 500,
	7, 90, 2, 2, 500, 88, 3, 2, 2, 2, 501, 502, 7, 82, 2, 2, 502, 503, 7, 84,
	2, 2, 503, 504, 7, 75, 2, 2, 504, 505, 7, 79, 2, 2, 505, 506, 7, 67, 2,
	2, 506, 507, 7, 84, 2, 2, 507, 508, 7, 91, 2, 2, 508, 90, 3, 2, 2, 2, 509,
	510, 7, 84, 2, 2, 510, 511, 7, 71, 2, 2, 511, 512, 7, 72, 2, 2, 512, 513,
	7, 71, 2, 2, 513, 514, 7, 84, 2, 2, 514, 515, 7, 71, 2, 2, 515, 516, 7,
	80, 2, 2, 516, 517, 7, 69, 2, 2, 517, 518, 7, 71, 2, 2, 518, 519, 7, 85,
	2, 2, 519, 92, 3, 2, 2, 2, 520, 521, 7, 95, 2, 2, 521, 94, 3, 2, 2, 2,
	522, 523, 7, 43, 2, 2, 523, 96, 3, 2, 2, 2, 524, 525, 7, 85, 2, 2, 525,
	526, 7, 74, 2, 2, 526, 527, 7, 69, 2, 2, 527, 528, 7, 71, 2, 2, 528, 529,
	7, 79, 2, 2, 529, 530, 7, 67, 2, 2, 530, 98, 3, 2, 2, 2, 531, 532, 7, 85,
	2, 2, 532, 533, 7, 71, 2, 2, 533, 534, 7, 86, 2, 2, 534, 100, 3, 2, 2,
	2, 535, 536, 7, 85, 2, 2, 536, 537, 7, 75, 2, 2, 537, 538, 7, 70, 2, 2,
	538, 102, 3, 2, 2, 2, 539, 540, 7, 85, 2, 2, 540, 541, 7, 86, 2, 2, 541,
	542, 7, 67, 2, 2, 542, 543, 7, 86, 2, 2, 543, 544, 7, 75, 2, 2, 544, 545,
	7, 85, 2, 2, 545, 546, 7, 86, 2, 2, 546, 547, 7, 75, 2, 2, 547, 548, 7,
	69, 2, 2, 548, 549, 7, 85, 2, 2, 549, 550, 7, 97, 2, 2, 550, 551, 7, 75,
	2, 2, 551, 552, 7, 80, 2, 2, 552, 553, 7, 69, 2, 2, 553, 554, 7, 84, 2,
	2, 554, 555, 7, 71, 2, 2, 555, 556, 7, 79, 2, 2, 556, 557, 7, 71, 2, 2,
	557, 558, 7, 80, 2, 2, 558, 559, 7, 86, 2, 2, 559, 560, 7, 67, 2, 2, 560,
	561, 7, 78, 2, 2, 561, 104, 3, 2, 2, 2, 562, 563, 7, 85, 2, 2, 563, 564,
	7, 86, 2, 2, 564, 565, 7, 67, 2, 2, 565, 566, 7, 86, 2, 2, 566, 567, 7,
	75, 2, 2, 567, 568, 7, 85, 2, 2, 568, 569, 7, 86, 2, 2, 569, 570, 7, 75,
	2, 2, 570, 571, 7, 69, 2, 2, 571, 572, 7, 85, 2, 2, 572, 573, 7, 97, 2,
	2, 573, 574, 7, 80, 2, 2, 574, 575, 7, 81, 2, 2, 575, 576, 7, 84, 2, 2,
	576, 577, 7, 71, 2, 2, 577, 578, 7, 69, 2, 2, 578, 579, 7, 81, 2, 2, 579,
	580, 7, 79, 2, 2, 580, 581, 7, 82, 2, 2, 581, 582, 7, 87, 2, 2, 582, 583,
	7, 86, 2, 2, 583, 584, 7, 71, 2, 2, 584, 106, 3, 2, 2, 2, 585, 586, 7,
	86, 2, 2, 586, 587, 7, 67, 2, 2, 587, 588, 7, 68, 2, 2, 588, 589, 7, 78,
	2, 2, 589, 590, 7, 71, 2, 2, 590, 108, 3, 2, 2, 2, 591, 592, 7, 86, 2,
	2, 592, 593, 7, 71, 2, 2, 593, 594, 7, 90, 2, 2, 594, 595, 7, 86, 2, 2,
	595, 596, 7, 75, 2, 2, 596, 597, 7, 79, 2, 2, 597, 598, 7, 67, 2, 2, 598,
	599, 7, 73, 2, 2, 599, 600, 7, 71, 2, 2, 600, 601, 7, 97, 2, 2, 601, 602,
	7, 81, 2, 2, 602, 603, 7, 80, 2, 2, 603, 110, 3, 2, 2, 2, 604, 605, 7,
	87, 2, 2, 605, 606, 7, 80, 2, 2, 606, 607, 7, 75, 2, 2, 607, 608, 7, 83,
	2, 2, 608, 609, 7, 87, 2, 2, 609, 610, 7, 71, 2, 2, 610, 112, 3, 2, 2,
	2, 611, 612, 7, 87, 2, 2, 612, 613, 7, 85, 2, 2, 613, 614, 7, 71, 2, 2,
	614, 114, 3, 2, 2, 2, 615, 616, 7, 87, 2, 2, 616, 617, 7, 85, 2, 2, 617,
	618, 7, 71, 2, 2, 618, 619, 7, 84, 2, 2, 619, 116, 3, 2, 2, 2, 620, 621,
	7, 93, 2, 2, 621, 622, 7, 120, 2, 2, 622, 623, 7, 99, 2, 2, 623, 624, 7,
	116, 2, 2, 624, 625, 7, 101, 2, 2, 625, 626, 7, 106, 2, 2, 626, 627, 7,
	99, 2, 2, 627, 628, 7, 116, 2, 2, 628, 629, 7, 95, 2, 2, 629, 118, 3, 2,
	2, 2, 630, 631, 7, 89, 2, 2, 631, 632, 7, 75, 2, 2, 632, 633, 7, 86, 2,
	2, 633, 634, 7, 74, 2, 2, 634, 120, 3, 2, 2, 2, 635, 636, 7, 93, 2, 2,
	636, 640, 5, 123, 62, 2, 637, 639, 5, 139, 70, 2, 638, 637, 3, 2, 2, 2,
	639, 642, 3, 2, 2, 2, 640, 638, 3, 2, 2, 2, 640, 641, 3, 2, 2, 2, 641,
	644, 3, 2, 2, 2, 642, 640, 3, 2, 2, 2, 643, 645, 5, 123, 62, 2, 644, 643,
	3, 2, 2, 2, 644, 645, 3, 2, 2, 2, 645, 646, 3, 2, 2, 2, 646, 647, 7, 95,
	2, 2, 647, 122, 3, 2, 2, 2, 648, 652, 5, 127, 64, 2, 649, 651, 5, 125,
	63, 2, 650, 649, 3, 2, 2, 2, 651, 654, 3, 2, 2, 2, 652, 650, 3, 2, 2, 2,
	652, 653, 3, 2, 2, 2, 653, 124, 3, 2, 2, 2, 654, 652, 3, 2, 2, 2, 655,
	658, 5, 129, 65, 2, 656, 658, 5, 131, 66, 2, 657, 655, 3, 2, 2, 2, 657,
	656, 3, 2, 2, 2, 658, 126, 3, 2, 2, 2, 659, 660, 9, 2, 2, 2, 660, 128,
	3, 2, 2, 2, 661, 662, 9, 3, 2, 2, 662, 130, 3, 2, 2, 2, 663, 664, 9, 4,
	2, 2, 664, 132, 3, 2, 2, 2, 665, 666, 9, 5, 2, 2, 666, 134, 3, 2, 2, 2,
	667, 668, 9, 6, 2, 2, 668, 136, 3, 2, 2, 2, 669, 670, 7, 94, 2, 2, 670,
	671, 9, 7, 2, 2, 671, 138, 3, 2, 2, 2, 672, 673, 9, 8, 2, 2, 673, 140,
	3, 2, 2, 2, 674, 675, 7, 73, 2, 2, 675, 676, 7, 81, 2, 2, 676, 142, 3,
	2, 2, 2, 677, 678, 5, 141, 71, 2, 678, 679, 3, 2, 2, 2, 679, 680, 8, 72,
	2, 2, 680, 144, 3, 2, 2, 2, 681, 683, 9, 9, 2, 2, 682, 681, 3, 2, 2, 2,
	683, 684, 3, 2, 2, 2, 684, 682, 3, 2, 2, 2, 684, 685, 3, 2, 2, 2, 685,
	686, 3, 2, 2, 2, 686, 687, 8, 73, 2, 2, 687, 146, 3, 2, 2, 2, 688, 690,
	7, 15, 2, 2, 689, 691, 7, 12, 2, 2, 690, 689, 3, 2, 2, 2, 690, 691, 3,
	2, 2, 2, 691, 694, 3, 2, 2, 2, 692, 694, 7, 12, 2, 2, 693, 688, 3, 2, 2,
	2, 693, 692, 3, 2, 2, 2, 694, 695, 3, 2, 2, 2, 695, 696, 8, 74, 2, 2, 696,
	148, 3, 2, 2, 2, 697, 698, 7, 49, 2, 2, 698, 699, 7, 44, 2, 2, 699, 703,
	3, 2, 2, 2, 700, 702, 11, 2, 2, 2, 701, 700, 3, 2, 2, 2, 702, 705, 3, 2,
	2, 2, 703, 704, 3, 2, 2, 2, 703, 701, 3, 2, 2, 2, 704, 706, 3, 2, 2, 2,
	705, 703, 3, 2, 2, 2, 706, 707, 7, 44, 2, 2, 707, 708, 7, 49, 2, 2, 708,
	709, 3, 2, 2, 2, 709, 710, 8, 75, 2, 2, 710, 150, 3, 2, 2, 2, 12, 2, 214,
	640, 644, 652, 657, 684, 690, 693, 703, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'SCHEMA'", "','", "'[PRIMARY]'", "'\"defaul\"'", "'.'", "'NONE'",
	"';'", "'QUOTED_IDENTIFIER'", "'OFF'", "", "'ADD'", "'ALLOW_PAGE_LOCKS'",
	"'ALLOW_ROW_LOCKS'", "'ALTER'", "'ANSI_NULLS'", "'ASC'", "'CHECK'", "'[char]'",
	"'CLUSTERED'", "'CONSTRAINT'", "'CREATE'", "'default'", "'DEFAULT_LANGUAGE'",
	"'DEFAULT_SCHEMA'", "'DESC'", "'='", "'FOR'", "'FOREIGN'", "'[geography]'",
	"'IGNORE_DUP_KEY'", "'[int]'", "'KEY'", "'('", "'['", "'LOGIN'", "'[money]'",
	"'NOCHECK'", "'NULL'", "'NOT'", "'NONCLUSTERED'", "'ON'", "'OPTIMIZE_FOR_SEQUENTIAL_KEY'",
	"'PAD_INDEX'", "'PRIMARY'", "'REFERENCES'", "']'", "')'", "'SHCEMA'", "'SET'",
	"'SID'", "'STATISTICS_INCREMENTAL'", "'STATISTICS_NORECOMPUTE'", "'TABLE'",
	"'TEXTIMAGE_ON'", "'UNIQUE'", "'USE'", "'USER'", "'[varchar]'", "'WITH'",
	"", "", "'GO'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "DigitSequence", "Add", "AllowPageLocks",
	"AllowRowLocks", "Alter", "AnsiNulls", "Asc", "Check", "CharType", "Clustered",
	"Constraint", "Create", "Default", "DefaultLanguage", "DefaultSchema",
	"Desc", "Equal", "For", "Foreign", "GeographyType", "IgnoreDupKey", "IntType",
	"Key", "LeftParen", "LeftBracket", "Login", "MoneyType", "NoCheck", "Null",
	"Not", "NonClustered", "On", "OptimizeForSequentialKey", "PadIndex", "Primary",
	"References", "RightBracket", "RightParen", "Schema", "Set", "Sid", "StatisticsIncremental",
	"StatisticsNoreCompute", "Table", "TextImageOn", "Unique", "Use", "User",
	"VarcharType", "With", "Identifier", "CharacterPart", "GoCommand", "IgnoredQuery",
	"Whitespace", "Newline", "BlockComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"DigitSequence", "Add", "AllowPageLocks", "AllowRowLocks", "Alter", "AnsiNulls",
	"Asc", "Check", "CharType", "Clustered", "Constraint", "Create", "Default",
	"DefaultLanguage", "DefaultSchema", "Desc", "Equal", "For", "Foreign",
	"GeographyType", "IgnoreDupKey", "IntType", "Key", "LeftParen", "LeftBracket",
	"Login", "MoneyType", "NoCheck", "Null", "Not", "NonClustered", "On", "OptimizeForSequentialKey",
	"PadIndex", "Primary", "References", "RightBracket", "RightParen", "Schema",
	"Set", "Sid", "StatisticsIncremental", "StatisticsNoreCompute", "Table",
	"TextImageOn", "Unique", "Use", "User", "VarcharType", "With", "Identifier",
	"CharacterSequence", "CharacterPart", "CharacterNoDigit", "Character",
	"SpecialCharacter", "DigitNoZero", "Digit", "SimpleEscapeSequence", "SimpleSpace",
	"GoCommand", "IgnoredQuery", "Whitespace", "Newline", "BlockComment",
}

type TransactSQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewTransactSQLLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *TransactSQLLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewTransactSQLLexer(input antlr.CharStream) *TransactSQLLexer {
	l := new(TransactSQLLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "TransactSQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TransactSQLLexer tokens.
const (
	TransactSQLLexerT__0                     = 1
	TransactSQLLexerT__1                     = 2
	TransactSQLLexerT__2                     = 3
	TransactSQLLexerT__3                     = 4
	TransactSQLLexerT__4                     = 5
	TransactSQLLexerT__5                     = 6
	TransactSQLLexerT__6                     = 7
	TransactSQLLexerT__7                     = 8
	TransactSQLLexerT__8                     = 9
	TransactSQLLexerDigitSequence            = 10
	TransactSQLLexerAdd                      = 11
	TransactSQLLexerAllowPageLocks           = 12
	TransactSQLLexerAllowRowLocks            = 13
	TransactSQLLexerAlter                    = 14
	TransactSQLLexerAnsiNulls                = 15
	TransactSQLLexerAsc                      = 16
	TransactSQLLexerCheck                    = 17
	TransactSQLLexerCharType                 = 18
	TransactSQLLexerClustered                = 19
	TransactSQLLexerConstraint               = 20
	TransactSQLLexerCreate                   = 21
	TransactSQLLexerDefault                  = 22
	TransactSQLLexerDefaultLanguage          = 23
	TransactSQLLexerDefaultSchema            = 24
	TransactSQLLexerDesc                     = 25
	TransactSQLLexerEqual                    = 26
	TransactSQLLexerFor                      = 27
	TransactSQLLexerForeign                  = 28
	TransactSQLLexerGeographyType            = 29
	TransactSQLLexerIgnoreDupKey             = 30
	TransactSQLLexerIntType                  = 31
	TransactSQLLexerKey                      = 32
	TransactSQLLexerLeftParen                = 33
	TransactSQLLexerLeftBracket              = 34
	TransactSQLLexerLogin                    = 35
	TransactSQLLexerMoneyType                = 36
	TransactSQLLexerNoCheck                  = 37
	TransactSQLLexerNull                     = 38
	TransactSQLLexerNot                      = 39
	TransactSQLLexerNonClustered             = 40
	TransactSQLLexerOn                       = 41
	TransactSQLLexerOptimizeForSequentialKey = 42
	TransactSQLLexerPadIndex                 = 43
	TransactSQLLexerPrimary                  = 44
	TransactSQLLexerReferences               = 45
	TransactSQLLexerRightBracket             = 46
	TransactSQLLexerRightParen               = 47
	TransactSQLLexerSchema                   = 48
	TransactSQLLexerSet                      = 49
	TransactSQLLexerSid                      = 50
	TransactSQLLexerStatisticsIncremental    = 51
	TransactSQLLexerStatisticsNoreCompute    = 52
	TransactSQLLexerTable                    = 53
	TransactSQLLexerTextImageOn              = 54
	TransactSQLLexerUnique                   = 55
	TransactSQLLexerUse                      = 56
	TransactSQLLexerUser                     = 57
	TransactSQLLexerVarcharType              = 58
	TransactSQLLexerWith                     = 59
	TransactSQLLexerIdentifier               = 60
	TransactSQLLexerCharacterPart            = 61
	TransactSQLLexerGoCommand                = 62
	TransactSQLLexerIgnoredQuery             = 63
	TransactSQLLexerWhitespace               = 64
	TransactSQLLexerNewline                  = 65
	TransactSQLLexerBlockComment             = 66
)
