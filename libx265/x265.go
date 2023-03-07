package libx265

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/x265-go/libx265common"
)

/*****************************************************************************
 * Copyright (C) 2013-2020 MulticoreWare, Inc
 *
 * Authors: Steve Borho <steve@borho.org>
 *          Min Chen <chenm003@163.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02111, USA.
 *
 * This program is also available under a commercial proprietary license.
 * For more information, contact us at license @ x265.com.
 *****************************************************************************/

// #ifndef X265_H
// const X265_H
// #include <stdint.h>
// #include <stdio.h>
// #include <sys/stat.h>
// #include "x265_config.h"
// #ifdef __cplusplus
// extern "C" {
// #endif

// #if _MSC_VER
// #pragma warning(disable: 4201) // non-standard extension used (nameless struct/union)
// #endif

/* x265_encoder:
 *      opaque handler for encoder */
// typedef struct x265_encoder x265_encoder;
type X265Encoder struct {
}

/* x265_picyuv:
 *      opaque handler for PicYuv */
// typedef struct x265_picyuv x265_picyuv;
type X265Picyuv struct {
}

/* Application developers planning to link against a shared library version of
 * libx265 from a Microsoft Visual Studio or similar development environment
 * will need to define X265_API_IMPORTS before including this header.
 * This clause does not apply to MinGW, similar development environments, or non
 * Windows platforms. */
// #ifdef X265_API_IMPORTS
// const X265_API __declspec(dllimport)
// #else
// const X265_API
// #endif
type NalUnitType ffcommon.FEnum

const (
	NAL_UNIT_CODED_SLICE_TRAIL_N = 0 + iota
	NAL_UNIT_CODED_SLICE_TRAIL_R
	NAL_UNIT_CODED_SLICE_TSA_N
	NAL_UNIT_CODED_SLICE_TSA_R
	NAL_UNIT_CODED_SLICE_STSA_N
	NAL_UNIT_CODED_SLICE_STSA_R
	NAL_UNIT_CODED_SLICE_RADL_N
	NAL_UNIT_CODED_SLICE_RADL_R
	NAL_UNIT_CODED_SLICE_RASL_N
	NAL_UNIT_CODED_SLICE_RASL_R
)
const (
	NAL_UNIT_CODED_SLICE_BLA_W_LP = 16 + iota
	NAL_UNIT_CODED_SLICE_BLA_W_RADL
	NAL_UNIT_CODED_SLICE_BLA_N_LP
	NAL_UNIT_CODED_SLICE_IDR_W_RADL
	NAL_UNIT_CODED_SLICE_IDR_N_LP
	NAL_UNIT_CODED_SLICE_CRA
)
const (
	NAL_UNIT_VPS = 32 + iota
	NAL_UNIT_SPS
	NAL_UNIT_PPS
	NAL_UNIT_ACCESS_UNIT_DELIMITER
	NAL_UNIT_EOS
	NAL_UNIT_EOB
	NAL_UNIT_FILLER_DATA
	NAL_UNIT_PREFIX_SEI
	NAL_UNIT_SUFFIX_SEI
)
const NAL_UNIT_UNSPECIFIED = 62
const NAL_UNIT_INVALID = 64

/* The data within the payload is already NAL-encapsulated; the type is merely
 * in the struct for easy access by the calling application.  All data returned
 * in an x265_nal, including the data in payload, is no longer valid after the
 * next call to x265_encoder_encode.  Thus it must be used or copied before
 * calling x265_encoder_encode again. */
type X265Nal struct {
	Type      ffcommon.FUint32T /* NalUnitType */
	SizeBytes ffcommon.FUint32T /* size in bytes */
	Payload   *ffcommon.FUint8T
}

const X265_LOOKAHEAD_MAX = 250

type X265LookaheadData struct {
	PlannedSatd       [X265_LOOKAHEAD_MAX + 1]ffcommon.FInt64T
	VbvCost           *ffcommon.FUint32T
	IntraVbvCost      *ffcommon.FUint32T
	SatdForVbv        *ffcommon.FUint32T
	IntraSatdForVbv   *ffcommon.FUint32T
	Keyframe          ffcommon.FInt
	LastMiniGopBFrame ffcommon.FInt
	PlannedType       [X265_LOOKAHEAD_MAX + 1]ffcommon.FInt
	Dts               ffcommon.FInt64T
	ReorderedPts      ffcommon.FInt64T
}

type X265AnalysisValidate struct {
	MaxNumReferences    ffcommon.FInt
	AnalysisReuseLevel  ffcommon.FInt
	SourceWidth         ffcommon.FInt
	SourceHeight        ffcommon.FInt
	KeyframeMax         ffcommon.FInt
	KeyframeMin         ffcommon.FInt
	OpenGOP             ffcommon.FInt
	Bframes             ffcommon.FInt
	BPyramid            ffcommon.FInt
	MaxCUSize           ffcommon.FInt
	MinCUSize           ffcommon.FInt
	IntraRefresh        ffcommon.FInt
	LookaheadDepth      ffcommon.FInt
	ChunkStart          ffcommon.FInt
	ChunkEnd            ffcommon.FInt
	CuTree              ffcommon.FInt
	CtuDistortionRefine ffcommon.FInt
	RightOffset         ffcommon.FInt
	BottomOffset        ffcommon.FInt
	frameDuplication    ffcommon.FInt
}

/* Stores intra analysis data for a single frame. This struct needs better packing */
type X265AnalysisIntraData struct {
	Depth       *ffcommon.FUint8T
	Modes       *ffcommon.FUint8T
	PartSizes   ffcommon.FCharPStruct
	ChromaModes *ffcommon.FUint8T
	CuQPOff     *ffcommon.FInt8T
}

type X265AnalysisMV struct {
	// union{
	//     struct { int32_t x, y; };

	//     int64_t word;
	// };
	X, Y ffcommon.FInt32T
}

/* Stores inter analysis data for a single frame */
type X265AnalysisInterData struct {
	Ref       *ffcommon.FInt32T
	Depth     *ffcommon.FUint8T
	Modes     *ffcommon.FUint8T
	PartSize  *ffcommon.FUint8T
	MergeFlag *ffcommon.FUint8T
	InterDir  *ffcommon.FUint8T
	MvpIdx    [2]*ffcommon.FUint8T
	RefIdx    [2]*ffcommon.FInt8T
	Mv        [2]*X265AnalysisMV
	SadCost   *ffcommon.FUint64T
	CuQPOff   *ffcommon.FInt8T
}

type X265WeightParam struct {
	Log2WeightDenom ffcommon.FUint32T
	InputWeight     ffcommon.FInt
	InputOffset     ffcommon.FInt
	WtPresent       ffcommon.FInt
}

// #if X265_DEPTH < 10
// typedef uint32_t sse_t;
// #else
// typedef uint64_t sse_t;
// #endif

type SseT = ffcommon.FUint32T

const CTU_DISTORTION_OFF = 0
const CTU_DISTORTION_INTERNAL = 1
const CTU_DISTORTION_EXTERNAL = 2

type X265AnalysisDistortionData struct {
	CtuDistortion          *SseT
	ScaledDistortion       *ffcommon.FDouble
	AverageDistortion      ffcommon.FDouble
	SdDistortion           ffcommon.FDouble
	HighDistortionCtuCount ffcommon.FUint32T
	LowDistortionCtuCount  ffcommon.FUint32T
	Offset                 *ffcommon.FDouble
	Threshold              *ffcommon.FDouble
}

const MAX_NUM_REF = 16
const EDGE_BINS = 2
const MAX_HIST_BINS = 1024

/* Stores all analysis data for a single frame */
type X265AnalysisData struct {
	SatdCost          ffcommon.FInt64T
	frameRecordSize   ffcommon.FUint32T
	Poc               ffcommon.FUint32T
	SliceType         ffcommon.FUint32T
	NumCUsInFrame     ffcommon.FUint32T
	NumPartitions     ffcommon.FUint32T
	DepthBytes        ffcommon.FUint32T
	EdgeHist          [EDGE_BINS]ffcommon.FInt32T
	YuvHist           [3][MAX_HIST_BINS]ffcommon.FInt32T
	BScenecut         ffcommon.FInt
	Wt                *X265WeightParam
	InterData         *X265AnalysisInterData
	IntraData         *X265AnalysisInterData
	NumCuInHeight     ffcommon.FUint32T
	Lookahead         X265LookaheadData
	ModeFlag          [2]*ffcommon.FUint8T
	SaveParam         X265AnalysisValidate
	DistortionData    *X265AnalysisDistortionData
	frameBits         ffcommon.FUint64T
	List0POC          [MAX_NUM_REF]ffcommon.FInt
	List1POC          [MAX_NUM_REF]ffcommon.FInt
	TotalIntraPercent ffcommon.FDouble
}

/* cu statistics */
type X265CuStats struct {
	PercentSkipCu            [4]ffcommon.FDouble    // Percentage of skip cu in all depths
	PercentMergeCu           [4]ffcommon.FDouble    // Percentage of merge cu in all depths
	PercentIntraDistribution [4][3]ffcommon.FDouble // Percentage of DC, Planar, Angular intra modes in all depths
	PercentInterDistribution [4][3]ffcommon.FDouble // Percentage of 2Nx2N inter, rect and amp in all depths
	PercentIntraNxN          ffcommon.FDouble       // Percentage of 4x4 cu

	/* All the above values will add up to 100%. */
}

/* pu statistics */
type X265PuStats struct {
	PercentSkipPu  [4]ffcommon.FDouble    // Percentage of skip cu in all depths
	PercentIntraPu [4]ffcommon.FDouble    // Percentage of intra modes in all depths
	PercentAmpPu   [4]ffcommon.FDouble    // Percentage of amp modes in all depths
	PercentInterPu [4][3]ffcommon.FDouble // Percentage of inter 2nx2n, 2nxn and nx2n in all depths
	PercentMergePu [4][3]ffcommon.FDouble // Percentage of merge 2nx2n, 2nxn and nx2n in all depth
	PercentNxN     ffcommon.FDouble

	/* All the above values will add up to 100%. */
}

/* Frame level statistics */
type X265FrameStats struct {
	Qp                  ffcommon.FDouble
	RateFactor          ffcommon.FDouble
	PsnrY               ffcommon.FDouble
	PsnrU               ffcommon.FDouble
	PsnrV               ffcommon.FDouble
	Psnr                ffcommon.FDouble
	Ssim                ffcommon.FDouble
	DecideWaitTime      ffcommon.FDouble
	Row0WaitTime        ffcommon.FDouble
	WallTime            ffcommon.FDouble
	RefWaitWallTime     ffcommon.FDouble
	TotalCTUTime        ffcommon.FDouble
	StallTime           ffcommon.FDouble
	AvgWPP              ffcommon.FDouble
	AvgLumaDistortion   ffcommon.FDouble
	AvgChromaDistortion ffcommon.FDouble
	AvgPsyEnergy        ffcommon.FDouble
	AvgResEnergy        ffcommon.FDouble
	AvgLumaLevel        ffcommon.FDouble
	BufferFill          ffcommon.FDouble
	Bits                ffcommon.FUint64T
	EncoderOrder        ffcommon.FInt
	Poc                 ffcommon.FInt
	CountRowBlocks      ffcommon.FInt
	List0POC            [MAX_NUM_REF]ffcommon.FInt
	List1POC            [MAX_NUM_REF]ffcommon.FInt
	MaxLumaLevel        ffcommon.FUint16T
	MinLumaLevel        ffcommon.FUint16T

	MaxChromaULevel ffcommon.FUint16T
	MinChromaULevel ffcommon.FUint16T
	AvgChromaULevel ffcommon.FDouble

	MaxChromaVLevel ffcommon.FUint16T
	MinChromaVLevel ffcommon.FUint16T
	AvgChromaVLevel ffcommon.FDouble

	SliceType                ffcommon.FChar
	BScenecut                ffcommon.FInt
	IpCostRatio              ffcommon.FDouble
	FrameLatency             ffcommon.FInt
	CuStats                  X265CuStats
	PuStats                  X265PuStats
	TotalFrameTime           ffcommon.FDouble
	VmafFrameScore           ffcommon.FDouble
	BufferFillFinal          ffcommon.FDouble
	UnclippedBufferFillFinal ffcommon.FDouble
	TLayer                   ffcommon.FUint64T
}

type X265CtuInfoT struct {
	CtuAddress    ffcommon.FInt32T
	CtuPartitions [64]ffcommon.FInt32T
	CtuInfo       ffcommon.FVoidP
}

type CTUInfo ffcommon.FEnum

const (
	NO_CTU_INFO     = 0
	HAS_CTU_INFO    = 1
	CTU_INFO_CHANGE = 2
)

type AnalysisRefineType ffcommon.FEnum

const (
	DEFAULT   = 0
	AVC_INFO  = 1
	HEVC_INFO = 2
)

/* Arbitrary User SEI
 * Payload size is in bytes and the payload pointer must be non-NULL.
 * Payload types and syntax can be found in Annex D of the H.265 Specification.
 * SEI Payload Alignment bits as described in Annex D must be included at the
 * end of the payload if needed. The payload should not be NAL-encapsulated.
 * Payloads are written in the order of input */
type SEIPayloadType ffcommon.FEnum

const (
	BUFFERING_PERIOD                     = 0
	PICTURE_TIMING                       = 1
	PAN_SCAN_RECT                        = 2
	FILLER_PAYLOAD                       = 3
	USER_DATA_REGISTERED_ITU_T_T35       = 4
	USER_DATA_UNREGISTERED               = 5
	RECOVERY_POINT                       = 6
	SCENE_INFO                           = 9
	FULL_FRAME_SNAPSHOT                  = 15
	PROGRESSIVE_REFINEMENT_SEGMENT_START = 16
	PROGRESSIVE_REFINEMENT_SEGMENT_END   = 17
	FILM_GRAIN_CHARACTERISTICS           = 19
	POST_FILTER_HINT                     = 22
	TONE_MAPPING_INFO                    = 23
	FRAME_PACKING                        = 45
	DISPLAY_ORIENTATION                  = 47
	SOP_DESCRIPTION                      = 128
	ACTIVE_PARAMETER_SETS                = 129
	DECODING_UNIT_INFO                   = 130
	TEMPORAL_LEVEL0_INDEX                = 131
	DECODED_PICTURE_HASH                 = 132
	SCALABLE_NESTING                     = 133
	REGION_REFRESH_INFO                  = 134
	MASTERING_DISPLAY_INFO               = 137
	CONTENT_LIGHT_LEVEL_INFO             = 144
	ALTERNATIVE_TRANSFER_CHARACTERISTICS = 147
)

type X265SeiPayload struct {
	PayloadSize ffcommon.FInt
	PayloadType SEIPayloadType
	Payload     *ffcommon.FUint8T
}

type X265Sei struct {
	NumPayloads ffcommon.FInt
	Payloads    *X265SeiPayload
}

type X265DolbyVisionRpu struct {
	PayloadSize ffcommon.FInt
	Payload     *ffcommon.FUint8T
}

/* Used to pass pictures into the encoder, and to get picture data back out of
 * the encoder.  The input and output semantics are different */
type X265Picture struct {
	/* presentation time stamp: user-specified, returned on output */
	Pts ffcommon.FInt64T

	/* display time stamp: ignored on input, copied from reordered pts. Returned
	 * on output */
	Dts ffcommon.FInt64T

	/* force quantizer for != X265_QP_AUTO */
	/* The value provided on input is returned with the same picture (POC) on
	 * output */
	UserData ffcommon.FVoidP

	/* Must be specified on input pictures, the number of planes is determined
	 * by the colorSpace value */
	Planes [3]ffcommon.FVoidP

	/* Stride is the number of bytes between row starts */
	Stride [3]ffcommon.FInt

	/* Must be specified on input pictures. x265_picture_init() will set it to
	 * the encoder's internal bit depth, but this field must describe the depth
	 * of the input pictures. Must be between 8 and 16. Values larger than 8
	 * imply 16bits per input sample. If input bit depth is larger than the
	 * internal bit depth, the encoder will down-shift pixels. Input samples
	 * larger than 8bits will be masked to internal bit depth. On output the
	 * bitDepth will be the internal encoder bit depth */
	BitDepth ffcommon.FInt

	/* Must be specified on input pictures: X265_TYPE_AUTO or other.
	 * x265_picture_init() sets this to auto, returned on output */
	SliceType ffcommon.FInt

	/* Ignored on input, set to picture count, returned on output */
	Poc ffcommon.FInt

	/* Must be specified on input pictures: X265_CSP_I420 or other. It must
	 * match the internal color space of the encoder. x265_picture_init() will
	 * initialize this value to the internal color space */
	ColorSpace ffcommon.FInt

	/* Force the slice base QP for this picture within the encoder. Set to 0
	 * to allow the encoder to determine base QP */
	Forceqp ffcommon.FInt

	/* If param.analysisLoad and param.analysisSave are disabled, this field is
	 * ignored on input and output. Else the user must call x265_alloc_analysis_data()
	 * to allocate analysis buffers for every picture passed to the encoder.
	 *
	 * On input when param.analysisLoad is enabled and analysisData
	 * member pointers are valid, the encoder will use the data stored here to
	 * reduce encoder work.
	 *
	 * On output when param.analysisSave is enabled and analysisData
	 * member pointers are valid, the encoder will write output analysis into
	 * this data structure */
	AnalysisData X265AnalysisData

	/* An array of quantizer offsets to be applied to this image during encoding.
	 * These are added on top of the decisions made by rateControl.
	 * Adaptive quantization must be enabled to use this feature. These quantizer
	 * offsets should be given for each 16x16 block (8x8 block, when qg-size is 8).
	 * Behavior if quant offsets differ between encoding passes is undefined. */
	QuantOffsets *ffcommon.FFloat

	/* Frame level statistics */
	FrameData X265FrameStats

	/* User defined SEI */
	UserSEI X265Sei

	/* Ratecontrol statistics for collecting the ratecontrol information.
	 * It is not used for collecting the last pass ratecontrol data in
	 * multi pass ratecontrol mode. */
	RcData ffcommon.FVoidP

	Framesize ffcommon.FSizeT

	Height ffcommon.FInt

	// pts is reordered in the order of encoding.
	ReorderedPts ffcommon.FInt64T

	//Dolby Vision RPU metadata
	Rpu X265DolbyVisionRpu

	FieldNum ffcommon.FInt

	//SEI picture structure message
	PicStruct ffcommon.FUint32T

	Width ffcommon.FInt
}

type X265_ME_METHODS ffcommon.FEnum

const (
	X265_DIA_SEARCH = iota
	X265_HEX_SEARCH
	X265_UMH_SEARCH
	X265_STAR_SEARCH
	X265_SEA
	X265_FULL_SEARCH
)

/* CPU flags */

/* x86 */
const X265_CPU_MMX = (1 << 0)
const X265_CPU_MMX2 = (1 << 1) /* MMX2 aka MMXEXT aka ISSE */
const X265_CPU_MMXEXT = X265_CPU_MMX2
const X265_CPU_SSE = (1 << 2)
const X265_CPU_SSE2 = (1 << 3)
const X265_CPU_LZCNT = (1 << 4)
const X265_CPU_SSE3 = (1 << 5)
const X265_CPU_SSSE3 = (1 << 6)
const X265_CPU_SSE4 = (1 << 7)    /* SSE4.1 */
const X265_CPU_SSE42 = (1 << 8)   /* SSE4.2 */
const X265_CPU_AVX = (1 << 9)     /* Requires OS support even if YMM registers aren't used. */
const X265_CPU_XOP = (1 << 10)    /* AMD XOP */
const X265_CPU_FMA4 = (1 << 11)   /* AMD FMA4 */
const X265_CPU_FMA3 = (1 << 12)   /* Intel FMA3 */
const X265_CPU_BMI1 = (1 << 13)   /* BMI1 */
const X265_CPU_BMI2 = (1 << 14)   /* BMI2 */
const X265_CPU_AVX2 = (1 << 15)   /* AVX2 */
const X265_CPU_AVX512 = (1 << 16) /* AVX-512 {F, CD, BW, DQ, VL}, requires OS support */
/* x86 modifiers */
const X265_CPU_CACHELINE_32 = (1 << 17) /* avoid memory loads that span the border between two cachelines */
const X265_CPU_CACHELINE_64 = (1 << 18) /* 32/64 is the size of a cacheline in bytes */
const X265_CPU_SSE2_IS_SLOW = (1 << 19) /* avoid most SSE2 functions on Athlon64 */
const X265_CPU_SSE2_IS_FAST = (1 << 20) /* a few functions are only faster on Core2 and Phenom */
const X265_CPU_SLOW_SHUFFLE = (1 << 21) /* The Conroe has a slow shuffle unit (relative to overall SSE performance) */
const X265_CPU_STACK_MOD4 = (1 << 22)   /* if stack is only mod4 and not mod16 */
const X265_CPU_SLOW_ATOM = (1 << 23)    /* The Atom is terrible: slow SSE unaligned loads, slow
 * SIMD multiplies, slow SIMD variable shifts, slow pshufb,
 * cacheline split penalties -- gather everything here that
 * isn't shared by other CPUs to avoid making half a dozen
 * new SLOW flags. */
const X265_CPU_SLOW_PSHUFB = (1 << 24)  /* such as on the Intel Atom */
const X265_CPU_SLOW_PALIGNR = (1 << 25) /* such as on the AMD Bobcat */

/* ARM */
const X265_CPU_ARMV6 = 0x0000001
const X265_CPU_NEON = 0x0000002          /* ARM NEON */
const X265_CPU_FAST_NEON_MRC = 0x0000004 /* Transfer from NEON to ARM register is fast (Cortex-A9) */

/* IBM Power8 */
const X265_CPU_ALTIVEC = 0x0000001

const X265_MAX_SUBPEL_LEVEL = 7

/* Log level */
const X265_LOG_NONE = (-1)
const X265_LOG_ERROR = 0
const X265_LOG_WARNING = 1
const X265_LOG_INFO = 2
const X265_LOG_DEBUG = 3
const X265_LOG_FULL = 4

const X265_B_ADAPT_NONE = 0
const X265_B_ADAPT_FAST = 1
const X265_B_ADAPT_TRELLIS = 2

const X265_REF_LIMIT_DEPTH = 1
const X265_REF_LIMIT_CU = 2

const X265_TU_LIMIT_BFS = 1
const X265_TU_LIMIT_DFS = 2
const X265_TU_LIMIT_NEIGH = 4

const X265_BFRAME_MAX = 16
const X265_MAX_FRAME_THREADS = 16

const X265_TYPE_AUTO = 0x0000 /* Let x265 choose the right type */
const X265_TYPE_IDR = 0x0001
const X265_TYPE_I = 0x0002
const X265_TYPE_P = 0x0003
const X265_TYPE_BREF = 0x0004 /* Non-disposable B-frame */
const X265_TYPE_B = 0x0005

// const IS_X265_TYPE_I(x) ((x) == X265_TYPE_I || (x) == X265_TYPE_IDR)
func IS_X265_TYPE_I(x ffcommon.FInt) bool {
	return x == X265_TYPE_I || x == X265_TYPE_IDR
}

// const IS_X265_TYPE_B(x) ((x) == X265_TYPE_B || (x) == X265_TYPE_BREF)
func IS_X265_TYPE_B(x ffcommon.FInt) bool {
	return x == X265_TYPE_B || x == X265_TYPE_BREF
}

const X265_QP_AUTO = 0

const X265_AQ_NONE = 0
const X265_AQ_VARIANCE = 1
const X265_AQ_AUTO_VARIANCE = 2
const X265_AQ_AUTO_VARIANCE_BIASED = 3
const X265_AQ_EDGE = 4
const x265_ADAPT_RD_STRENGTH = 4
const X265_REFINE_INTER_LEVELS = 3

/* NOTE! For this release only X265_CSP_I420 and X265_CSP_I444 are supported */
/* Supported internal color space types (according to semantics of chroma_format_idc) */
const X265_CSP_I400 = 0  /* yuv 4:0:0 planar */
const X265_CSP_I420 = 1  /* yuv 4:2:0 planar */
const X265_CSP_I422 = 2  /* yuv 4:2:2 planar */
const X265_CSP_I444 = 3  /* yuv 4:4:4 planar */
const X265_CSP_COUNT = 4 /* Number of supported internal color spaces */

/* These color spaces will eventually be supported as input pictures. The pictures will
 * be converted to the appropriate planar color spaces at ingest */
const X265_CSP_NV12 = 4 /* yuv 4:2:0, with one y plane and one packed u+v */
const X265_CSP_NV16 = 5 /* yuv 4:2:2, with one y plane and one packed u+v */

/* Interleaved color-spaces may eventually be supported as input pictures */
const X265_CSP_BGR = 6        /* packed bgr 24bits   */
const X265_CSP_BGRA = 7       /* packed bgr 32bits   */
const X265_CSP_RGB = 8        /* packed rgb 24bits   */
const X265_CSP_MAX = 9        /* end of list */
const X265_EXTENDED_SAR = 255 /* aspect ratio explicitly specified as width:height */
/* Analysis options */
const X265_ANALYSIS_OFF = 0
const X265_ANALYSIS_SAVE = 1
const X265_ANALYSIS_LOAD = 2

const FORWARD = 1
const BACKWARD = 2
const BI_DIRECTIONAL = 3
const SLICE_TYPE_DELTA = 0.3 /* The offset decremented or incremented for P-frames or b-frames respectively*/
const BACKWARD_WINDOW = 1    /* Scenecut window before a scenecut */
const FORWARD_WINDOW = 2     /* Scenecut window after a scenecut */
const BWD_WINDOW_DELTA = 0.4

const X265_MAX_GOP_CONFIG = 3
const X265_MAX_GOP_LENGTH = 16
const MAX_T_LAYERS = 7

type X265CliCsp struct {
	Planes ffcommon.FInt
	Width  [3]ffcommon.FInt
	Height [3]ffcommon.FInt
}

// static const x265_cli_csp x265_cli_csps[] =
// {
//     { 1, { 0, 0, 0 }, { 0, 0, 0 } }, /* i400 */
//     { 3, { 0, 1, 1 }, { 0, 1, 1 } }, /* i420 */
//     { 3, { 0, 1, 1 }, { 0, 0, 0 } }, /* i422 */
//     { 3, { 0, 0, 0 }, { 0, 0, 0 } }, /* i444 */
//     { 2, { 0, 0 },    { 0, 1 } },    /* nv12 */
//     { 2, { 0, 0 },    { 0, 0 } },    /* nv16 */
// };

var X265CliCsps = [...]X265CliCsp{
	X265CliCsp{1, [3]ffcommon.FInt{0, 0, 0}, [3]ffcommon.FInt{0, 0, 0}}, /* i400 */
	X265CliCsp{3, [3]ffcommon.FInt{0, 1, 1}, [3]ffcommon.FInt{0, 1, 1}}, /* i420 */
	X265CliCsp{3, [3]ffcommon.FInt{0, 1, 1}, [3]ffcommon.FInt{0, 0, 0}}, /* i422 */
	X265CliCsp{3, [3]ffcommon.FInt{0, 0, 0}, [3]ffcommon.FInt{0, 0, 0}}, /* i444 */
	X265CliCsp{2, [3]ffcommon.FInt{0, 0}, [3]ffcommon.FInt{0, 1}},       /* nv12 */
	X265CliCsp{2, [3]ffcommon.FInt{0, 0}, [3]ffcommon.FInt{0, 0}},       /* nv16 */
}

/* rate tolerance method */
type X265_RC_METHODS ffcommon.FEnum

const (
	X265_RC_ABR = iota
	X265_RC_CQP
	X265_RC_CRF
)

/* slice type statistics */
type X265SliceTypeStats struct {
	AvgQp   ffcommon.FDouble
	Bitrate ffcommon.FDouble
	PsnrY   ffcommon.FDouble
	PsnrU   ffcommon.FDouble
	PsnrV   ffcommon.FDouble
	Ssim    ffcommon.FDouble
	NumPics ffcommon.FUint32T
}

/* Output statistics from encoder */
type X265Stats struct {
	GlobalPsnrY         ffcommon.FDouble
	GlobalPsnrU         ffcommon.FDouble
	GlobalPsnrV         ffcommon.FDouble
	GlobalPsnr          ffcommon.FDouble
	GlobalSsim          ffcommon.FDouble
	ElapsedEncodeTime   ffcommon.FDouble   /* wall time since encoder was opened */
	ElapsedVideoTime    ffcommon.FDouble   /* encoded picture count / frame rate */
	Bitrate             ffcommon.FDouble   /* accBits / elapsed video time */
	AggregateVmafScore  ffcommon.FDouble   /* aggregate VMAF score for input video*/
	AccBits             ffcommon.FUint64T  /* total bits output thus far */
	EncodedPictureCount ffcommon.FUint32T  /* number of output pictures thus far */
	TotalWPFrames       ffcommon.FUint32T  /* number of uni-directional weighted frames used */
	StatsI              X265SliceTypeStats /* statistics of I slice */
	StatsP              X265SliceTypeStats /* statistics of P slice */
	StatsB              X265SliceTypeStats /* statistics of B slice */
	MaxCLL              ffcommon.FUint16T  /* maximum content light level */
	MaxFALL             ffcommon.FUint16T  /* maximum frame average light level */
}

// /* String values accepted by x265_param_parse() (and CLI) for various parameters */
// static const char * const x265_motion_est_names[] = { "dia", "hex", "umh", "star", "sea", "full", 0 };
var X265MotionEstNames = []string{"dia", "hex", "umh", "star", "sea", "full"}

// static const char * const x265_source_csp_names[] = { "i400", "i420", "i422", "i444", "nv12", "nv16", 0 };
var X265SourceCspNames = []string{"i400", "i420", "i422", "i444", "nv12", "nv16"}

// static const char * const x265_video_format_names[] = { "component", "pal", "ntsc", "secam", "mac", "unknown", 0 };
var X265VideoFormatNames = []string{"component", "pal", "ntsc", "secam", "mac", "unknown"}

// static const char * const x265_fullrange_names[] = { "limited", "full", 0 };
var X265FullrangeNames = []string{"limited", "full"}

// static const char * const x265_colorprim_names[] = { "reserved", "bt709", "unknown", "reserved", "bt470m", "bt470bg", "smpte170m", "smpte240m", "film", "bt2020", "smpte428", "smpte431", "smpte432", 0 };
var X265ColorprimNames = []string{"reserved", "bt709", "unknown", "reserved", "bt470m", "bt470bg", "smpte170m", "smpte240m", "film", "bt2020", "smpte428", "smpte431", "smpte432"}

// static const char * const x265_transfer_names[] = { "reserved", "bt709", "unknown", "reserved", "bt470m", "bt470bg", "smpte170m", "smpte240m", "linear", "log100",
//
//	"log316", "iec61966-2-4", "bt1361e", "iec61966-2-1", "bt2020-10", "bt2020-12",
//	"smpte2084", "smpte428", "arib-std-b67", 0 };
var X265TransferNames = []string{"reserved", "bt709", "unknown", "reserved", "bt470m", "bt470bg", "smpte170m", "smpte240m", "linear", "log100", "log316", "iec61966-2-4", "bt1361e", "iec61966-2-1", "bt2020-10", "bt2020-12", "smpte2084", "smpte428", "arib-std-b67"}

// static const char * const x265_colmatrix_names[] = { "gbr", "bt709", "unknown", "", "fcc", "bt470bg", "smpte170m", "smpte240m",
//
//	"ycgco", "bt2020nc", "bt2020c", "smpte2085", "chroma-derived-nc", "chroma-derived-c", "ictcp", 0 };
var X265ColmatrixNames = []string{"gbr", "bt709", "unknown", "", "fcc", "bt470bg", "smpte170m", "smpte240m", "ycgco", "bt2020nc", "bt2020c", "smpte2085", "chroma-derived-nc", "chroma-derived-c", "ictcp"}

// static const char * const x265_sar_names[] = { "unknown", "1:1", "12:11", "10:11", "16:11", "40:33", "24:11", "20:11",
//
//	"32:11", "80:33", "18:11", "15:11", "64:33", "160:99", "4:3", "3:2", "2:1", 0 };
var X265SarNames = []string{"unknown", "1:1", "12:11", "10:11", "16:11", "40:33", "24:11", "20:11", "32:11", "80:33", "18:11", "15:11", "64:33", "160:99", "4:3", "3:2", "2:1"}

// static const char * const x265_interlace_names[] = { "prog", "tff", "bff", 0 };
var X265InterlaceNames = []string{"prog", "tff", "bff"}

// static const char * const x265_analysis_names[] = { "off", "save", "load", 0 };
var X265AnalysisNames = []string{"off", "save", "load"}

// struct x265_zone;
// struct x265_param;
// type X265Param struct{}

// /* Zones: override ratecontrol for specific sections of the video.
//   - If zones overlap, whichever comes later in the list takes precedence. */
type X265Zone struct {
	StartFrame, endFrame ffcommon.FInt /* range of frame numbers */
	BForceQp             ffcommon.FInt /* whether to use qp vs bitrate factor */
	Qp                   ffcommon.FInt
	BitrateFactor        ffcommon.FFloat
	ZoneParam            *X265Param
	RelativeComplexity   ffcommon.FDouble
}
type FILE struct {
}

/* data to calculate aggregate VMAF score */
type X265VmafData struct {
	Width            ffcommon.FInt
	Height           ffcommon.FInt
	Offset           ffcommon.FSizeT
	InternalBitDepth ffcommon.FInt
	Reference_file   *FILE /* FILE pointer for input file */
	Distorted_file   *FILE /* FILE pointer for recon file generated*/
}

/* data to calculate frame level VMAF score */
type X265VmafFramedata struct {
	Width            ffcommon.FInt
	Height           ffcommon.FInt
	Frame_set        ffcommon.FInt
	InternalBitDepth ffcommon.FInt
	Reference_frame  ffcommon.FVoidP /* points to fenc of particular frame */
	Distorted_frame  ffcommon.FVoidP /* points to recon of particular frame */
}

/* common data needed to calculate both frame level and video level VMAF scores */
type X265VmafCommondata struct {
	Format               ffcommon.FCharPStruct
	Model_path           ffcommon.FCharPStruct
	Log_path             ffcommon.FCharPStruct
	Log_fmt              ffcommon.FCharPStruct
	Disable_clip         ffcommon.FInt
	Disable_avx          ffcommon.FInt
	Enable_transform     ffcommon.FInt
	Phone_model          ffcommon.FInt
	Psnr                 ffcommon.FInt
	Ssim                 ffcommon.FInt
	Ms_ssim              ffcommon.FInt
	Pool                 ffcommon.FCharPStruct
	Thread               ffcommon.FInt
	Subsample            ffcommon.FInt
	Enable_conf_interval ffcommon.FInt
}

// static const x265_vmaf_commondata vcd[] = { { NULL, (char *)"/usr/local/share/model/vmaf_v0.6.1.pkl", NULL, NULL, 0, 0, 0, 0, 0, 0, 0, NULL, 0, 1, 0 } };
// var vcd = [...]X265VmafCommondata{X265VmafCommondata{uintptr(0), ffcommon.BytePtrFromString("/usr/local/share/model/vmaf_v0.6.1.pkl"), uintptr(0), uintptr(0), 0, 0, 0, 0, 0, 0, 0, uintptr(0), 0, 1, 0}}

type X265TemporalLayer struct {
	Poc_offset ffcommon.FInt   /* POC offset */
	Layer      ffcommon.FInt8T /* Current layer */
	Qp_offset  ffcommon.FInt8T /* QP offset */
}

// static const int8_t x265_temporal_layer_bframes[MAX_T_LAYERS] = {-1, -1, 3, 7, 15, -1, -1};
// var X265TemporalLayerBframes = [MAX_T_LAYERS]ffcommon.FInt8T{-1, -1, 3, 7, 15, -1, -1}

// static const int8_t x265_gop_ra_length[X265_MAX_GOP_CONFIG] = { 4, 8, 16};
var X265GopRaLength = [X265_MAX_GOP_CONFIG]ffcommon.FInt8T{4, 8, 16}

var x265_gop_ra = [X265_MAX_GOP_CONFIG][X265_MAX_GOP_LENGTH]X265TemporalLayer{
	{
		{
			4,
			0,
			1,
		},
		{
			2,
			1,
			5,
		},
		{
			1,
			2,
			3,
		},
		{
			3,
			2,
			5,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
	},

	{
		{
			8,
			0,
			1,
		},
		{
			4,
			1,
			5,
		},
		{
			2,
			2,
			4,
		},
		{
			1,
			3,
			5,
		},
		{
			3,
			3,
			2,
		},
		{
			6,
			2,
			5,
		},
		{
			5,
			3,
			4,
		},
		{
			7,
			3,
			5,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
		{
			-1,
			-1,
			-1,
		},
	},
	{
		{
			16,
			0,
			1,
		},
		{
			8,
			1,
			6,
		},
		{
			4,
			2,
			5,
		},
		{
			2,
			3,
			6,
		},
		{
			1,
			4,
			4,
		},
		{
			3,
			4,
			6,
		},
		{
			6,
			3,
			5,
		},
		{
			5,
			4,
			6,
		},
		{
			7,
			4,
			1,
		},
		{
			12,
			2,
			6,
		},
		{
			10,
			3,
			5,
		},
		{
			9,
			4,
			6,
		},
		{
			11,
			4,
			4,
		},
		{
			14,
			3,
			6,
		},
		{
			13,
			4,
			5,
		},
		{
			15,
			4,
			6,
		},
	},
}

type X265_DATA_SHARE_MODES ffcommon.FEnum

const (
	X265_SHARE_MODE_FILE = iota
	X265_SHARE_MODE_SHAREDMEM
)

/* x265 input parameters
 *
 * For version safety you may use x265_param_alloc/free() to manage the
 * allocation of x265_param instances, and x265_param_parse() to assign values
 * by name.  By never dereferencing param fields in your own code you can treat
 * x265_param as an opaque data structure */
type X265Param struct {
	/* x265_param_default() will auto-detect this cpu capability bitmap.  it is
	 * recommended to not change this value unless you know the cpu detection is
	 * somehow flawed on your target hardware. The asm function tables are
	 * process global, the first encoder configures them for all encoders */
	Cpuid ffcommon.FInt
	/*== Parallelism Features ==*/

	/* Number of concurrently encoded frames between 1 and X265_MAX_FRAME_THREADS
	 * or 0 for auto-detection. By default x265 will use a number of frame
	 * threads empirically determined to be optimal for your CPU core count,
	 * between 2 and 6.  Using more than one frame thread causes motion search
	 * in the down direction to be clamped but otherwise encode behavior is
	 * unaffected. With CQP rate control the output bitstream is deterministic
	 * for all values of frameNumThreads greater than 1. All other forms of
	 * rate-control can be negatively impacted by increases to the number of
	 * frame threads because the extra concurrency adds uncertainty to the
	 * bitrate estimations. Frame parallelism is generally limited by the the
	 * is generally limited by the the number of CU rows
	 *
	 * When thread pools are used, each frame thread is assigned to a single
	 * pool and the frame thread itself is given the node affinity of its pool.
	 * But when no thread pools are used no node affinity is assigned. */
	FrameNumThreads ffcommon.FInt

	/* Comma seperated list of threads per NUMA node. If "none", then no worker
	 * pools are created and only frame parallelism is possible. If NULL or ""
	 * (default) x265 will use all available threads on each NUMA node.
	 *
	 * '+'  is a special value indicating all cores detected on the node
	 * '*'  is a special value indicating all cores detected on the node and all
	 *      remaining nodes.
	 * '-'  is a special value indicating no cores on the node, same as '0'
	 *
	 * example strings for a 4-node system:
	 *   ""        - default, unspecified, all numa nodes are used for thread pools
	 *   "*"       - same as default
	 *   "none"    - no thread pools are created, only frame parallelism possible
	 *   "-"       - same as "none"
	 *   "10"      - allocate one pool, using up to 10 cores on node 0
	 *   "-,+"     - allocate one pool, using all cores on node 1
	 *   "+,-,+"   - allocate two pools, using all cores on nodes 0 and 2
	 *   "+,-,+,-" - allocate two pools, using all cores on nodes 0 and 2
	 *   "-,*"     - allocate three pools, using all cores on nodes 1, 2 and 3
	 *   "8,8,8,8" - allocate four pools with up to 8 threads in each pool
	 *
	 * The total number of threads will be determined by the number of threads
	 * assigned to all nodes. The worker threads will each be given affinity for
	 * their node, they will not be allowed to migrate between nodes, but they
	 * will be allowed to move between CPU cores within their node.
	 *
	 * If the three pool features: bEnableWavefront, bDistributeModeAnalysis and
	 * bDistributeMotionEstimation are all disabled, then numaPools is ignored
	 * and no thread pools are created.
	 *
	 * If "none" is specified, then all three of the thread pool features are
	 * implicitly disabled.
	 *
	 * Multiple thread pools will be allocated for any NUMA node with more than
	 * 64 logical CPU cores. But any given thread pool will always use at most
	 * one NUMA node.
	 *
	 * Frame encoders are distributed between the available thread pools, and
	 * the encoder will never generate more thread pools than frameNumThreads */
	NumaPools ffcommon.FCharPStruct

	/* Enable wavefront parallel processing, greatly increases parallelism for
	 * less than 1% compression efficiency loss. Requires a thread pool, enabled
	 * by default */
	BEnableWavefront ffcommon.FInt

	/* Use multiple threads to measure CU mode costs. Recommended for many core
	 * CPUs. On RD levels less than 5, it may not offload enough work to warrant
	 * the overhead. It is useful with the slow preset since it has the
	 * rectangular predictions enabled. At RD level 5 and 6 (preset slower and
	 * below), this feature should be an unambiguous win if you have CPU
	 * cores available for work. Default disabled */
	BDistributeModeAnalysis ffcommon.FInt

	/* Use multiple threads to perform motion estimation to (ME to one reference
	 * per thread). Recommended for many core CPUs. The more references the more
	 * motion searches there will be to distribute. This option is often not a
	 * win, particularly in video sequences with low motion. Default disabled */
	BDistributeMotionEstimation ffcommon.FInt

	/*== Logging Features ==*/

	/* Enable analysis and logging distribution of CUs. Now deprecated */
	BLogCuStats ffcommon.FInt

	/* Enable the measurement and reporting of PSNR. Default is enabled */
	BEnablePsnr ffcommon.FInt

	/* Enable the measurement and reporting of SSIM. Default is disabled */
	BEnableSsim ffcommon.FInt

	/* The level of logging detail emitted by the encoder. X265_LOG_NONE to
	 * X265_LOG_FULL, default is X265_LOG_INFO */
	LogLevel ffcommon.FInt

	/* Level of csv logging. 0 is summary, 1 is frame level logging,
	 * 2 is frame level logging with performance statistics */
	CsvLogLevel ffcommon.FInt

	/* filename of CSV log. If csvLogLevel is non-zero, the encoder will emit
	 * per-slice statistics to this log file in encode order. Otherwise the
	 * encoder will emit per-stream statistics into the log file when
	 * x265_encoder_log is called (presumably at the end of the encode) */
	Csvfn ffcommon.FCharPStruct

	/*== Internal Picture Specification ==*/

	/* Internal encoder bit depth. If x265 was compiled to use 8bit pixels
	 * (HIGH_BIT_DEPTH=0), this field must be 8, else this field must be 10.
	 * Future builds may support 12bit pixels. */
	InternalBitDepth ffcommon.FInt

	/* Color space of internal pictures, must match color space of input
	 * pictures */
	InternalCsp ffcommon.FInt

	/* Numerator and denominator of frame rate */
	FpsNum   ffcommon.FUint32T
	FpsDenom ffcommon.FUint32T

	/* Width (in pixels) of the source pictures. If this width is not an even
	 * multiple of 4, the encoder will pad the pictures internally to meet this
	 * minimum requirement. All valid HEVC widths are supported */
	SourceWidth ffcommon.FInt

	/* Height (in pixels) of the source pictures. If this height is not an even
	 * multiple of 4, the encoder will pad the pictures internally to meet this
	 * minimum requirement. All valid HEVC heights are supported */
	SourceHeight ffcommon.FInt

	/* Interlace type of source pictures. 0 - progressive pictures (default).
	 * 1 - top field first, 2 - bottom field first. HEVC encodes interlaced
	 * content as fields, they must be provided to the encoder in the correct
	 * temporal order */
	InterlaceMode ffcommon.FInt

	/* Total Number of frames to be encoded, calculated from the user input
	 * (--frames) and (--seek). In case, the input is read from a pipe, this can
	 * remain as 0. It is later used in 2 pass RateControl, hence storing the
	 * value in param */
	TotalFrames ffcommon.FInt

	/*== Profile / Tier / Level ==*/

	/* Note: the profile is specified by x265_param_apply_profile() */

	/* Minimum decoder requirement level. Defaults to 0, which implies auto-
	 * detection by the encoder. If specified, the encoder will attempt to bring
	 * the encode specifications within that specified level. If the encoder is
	 * unable to reach the level it issues a warning and emits the actual
	 * decoder requirement. If the requested requirement level is higher than
	 * the actual level, the actual requirement level is signaled. The value is
	 * an specified as an integer with the level times 10, for example level
	 * "5.1" is specified as 51, and level "5.0" is specified as 50. */
	LevelIdc ffcommon.FInt

	/* if levelIdc is specified (non-zero) this flag will differentiate between
	 * Main (0) and High (1) tier. Default is Main tier (0) */
	BHighTier ffcommon.FInt

	/* Enable UHD Blu-ray compatibility support. If specified, the encoder will
	 * attempt to modify/set the encode specifications. If the encoder is unable
	 * to do so, this option will be turned OFF. */
	UhdBluray ffcommon.FInt

	/* The maximum number of L0 references a P or B slice may use. This
	 * influences the size of the decoded picture buffer. The higher this
	 * number, the more reference frames there will be available for motion
	 * search, improving compression efficiency of most video at a cost of
	 * performance. Value must be between 1 and 16, default is 3 */
	MaxNumReferences ffcommon.FInt

	/* Allow libx265 to emit HEVC bitstreams which do not meet strict level
	 * requirements. Defaults to false */
	BAllowNonConformance ffcommon.FInt

	/*== Bitstream Options ==*/

	/* Flag indicating whether VPS, SPS and PPS headers should be output with
	 * each keyframe. Default false */
	BRepeatHeaders ffcommon.FInt

	/* Flag indicating whether the encoder should generate start codes (Annex B
	 * format) or length (file format) before NAL units. Default true, Annex B.
	 * Muxers should set this to the correct value */
	BAnnexB ffcommon.FInt

	/* Flag indicating whether the encoder should emit an Access Unit Delimiter
	 * NAL at the start of every access unit. Default false */
	BEnableAccessUnitDelimiters ffcommon.FInt

	/* Enables the buffering period SEI and picture timing SEI to signal the HRD
	 * parameters. Default is disabled */
	BEmitHRDSEI ffcommon.FInt

	/* Enables the emission of a user data SEI with the stream headers which
	 * describes the encoder version, build info, and parameters. This is
	 * very helpful for debugging, but may interfere with regression tests.
	 * Default enabled */
	BEmitInfoSEI ffcommon.FInt

	/* Enable the generation of SEI messages for each encoded frame containing
	 * the hashes of the three reconstructed picture planes. Most decoders will
	 * validate those hashes against the reconstructed images it generates and
	 * report any mismatches. This is essentially a debugging feature.  Hash
	 * types are MD5(1), CRC(2), Checksum(3).  Default is 0, none */
	DecodedPictureHashSEI ffcommon.FInt

	/* Enable Temporal Sub Layers while encoding, signals NAL units of coded
	 * slices with their temporalId. Output bitstreams can be extracted either
	 * at the base temporal layer (layer 0) with roughly half the frame rate or
	 * at a higher temporal layer (layer 1) that decodes all the frames in the
	 * sequence. */
	BEnableTemporalSubLayers ffcommon.FInt

	/*== GOP structure and slice type decisions (lookahead) ==*/

	/* Enable open GOP - meaning I slices are not necessarily IDR and thus frames
	 * encoded after an I slice may reference frames encoded prior to the I
	 * frame which have remained in the decoded picture buffer.  Open GOP
	 * generally has better compression efficiency and negligible encoder
	 * performance impact, but the use case may preclude it.  Default true */
	BOpenGOP ffcommon.FInt

	/* Scene cuts closer together than this are coded as I, not IDR. */
	KeyframeMin ffcommon.FInt

	/* Maximum keyframe distance or intra period in number of frames. If 0 or 1,
	 * all frames are I frames. A negative value is casted to MAX_INT internally
	 * which effectively makes frame 0 the only I frame. Default is 250 */
	KeyframeMax ffcommon.FInt

	/* Maximum consecutive B frames that can be emitted by the lookahead. When
	 * b-adapt is 0 and keyframMax is greater than bframes, the lookahead emits
	 * a fixed pattern of `bframes` B frames between each P.  With b-adapt 1 the
	 * lookahead ignores the value of bframes for the most part.  With b-adapt 2
	 * the value of bframes determines the search (POC) distance performed in
	 * both directions, quadratically increasing the compute load of the
	 * lookahead.  The higher the value, the more B frames the lookahead may
	 * possibly use consecutively, usually improving compression. Default is 3,
	 * maximum is 16 */
	Bframes ffcommon.FInt

	/* Sets the operating mode of the lookahead.  With b-adapt 0, the GOP
	 * structure is fixed based on the values of keyframeMax and bframes.
	 * With b-adapt 1 a light lookahead is used to chose B frame placement.
	 * With b-adapt 2 (trellis) a viterbi B path selection is performed */
	BFrameAdaptive ffcommon.FInt

	/* When enabled, the encoder will use the B frame in the middle of each
	 * mini-GOP larger than 2 B frames as a motion reference for the surrounding
	 * B frames.  This improves compression efficiency for a small performance
	 * penalty.  Referenced B frames are treated somewhere between a B and a P
	 * frame by rate control.  Default is enabled. */
	BBPyramid ffcommon.FInt

	/* A value which is added to the cost estimate of B frames in the lookahead.
	 * It may be a positive value (making B frames appear less expensive, which
	 * biases the lookahead to choose more B frames) or negative, which makes the
	 * lookahead choose more P frames. Default is 0, there are no limits */
	BFrameBias ffcommon.FInt

	/* The number of frames that must be queued in the lookahead before it may
	 * make slice decisions. Increasing this value directly increases the encode
	 * latency. The longer the queue the more optimally the lookahead may make
	 * slice decisions, particularly with b-adapt 2. When cu-tree is enabled,
	 * the length of the queue linearly increases the effectiveness of the
	 * cu-tree analysis. Default is 40 frames, maximum is 250 */
	LookaheadDepth ffcommon.FInt

	/* Use multiple worker threads to measure the estimated cost of each frame
	 * within the lookahead. When bFrameAdaptive is 2, most frame cost estimates
	 * will be performed in batch mode, many cost estimates at the same time,
	 * and lookaheadSlices is ignored for batched estimates. The effect on
	 * performance can be quite small.  The higher this parameter, the less
	 * accurate the frame costs will be (since context is lost across slice
	 * boundaries) which will result in less accurate B-frame and scene-cut
	 * decisions. Default is 0 - disabled. 1 is the same as 0. Max 16 */
	LookaheadSlices ffcommon.FInt

	/* An arbitrary threshold which determines how aggressively the lookahead
	 * should detect scene cuts for cost based scenecut detection.
	 * The default (40) is recommended. */
	ScenecutThreshold ffcommon.FInt

	/* Replace keyframes by using a column of intra blocks that move across the video
	 * from one side to the other, thereby "refreshing" the image. In effect, instead of a
	 * big keyframe, the keyframe is "spread" over many frames. */
	BIntraRefresh ffcommon.FInt

	/*== Coding Unit (CU) definitions ==*/

	/* Maximum CU width and height in pixels.  The size must be 64, 32, or 16.
	 * The higher the size, the more efficiently x265 can encode areas of low
	 * complexity, greatly improving compression efficiency at large
	 * resolutions.  The smaller the size, the more effective wavefront and
	 * frame parallelism will become because of the increase in rows. default 64
	 * All encoders within the same process must use the same maxCUSize, until
	 * all encoders are closed and x265_cleanup() is called to reset the value. */
	MaxCUSize ffcommon.FUint32T

	/* Minimum CU width and height in pixels.  The size must be 64, 32, 16, or
	 * 8. Default 8. All encoders within the same process must use the same
	 * minCUSize. */
	MinCUSize ffcommon.FUint32T

	/* Enable rectangular motion prediction partitions (vertical and
	 * horizontal), available at all CU depths from 64x64 to 8x8. Default is
	 * disabled */
	BEnableRectInter ffcommon.FInt

	/* Enable asymmetrical motion predictions.  At CU depths 64, 32, and 16, it
	 * is possible to use 25%/75% split partitions in the up, down, right, left
	 * directions. For some material this can improve compression efficiency at
	 * the cost of extra analysis. bEnableRectInter must be enabled for this
	 * feature to be used. Default disabled */
	BEnableAMP ffcommon.FInt

	/*== Residual Quadtree Transform Unit (TU) definitions ==*/

	/* Maximum TU width and height in pixels.  The size must be 32, 16, 8 or 4.
	 * The larger the size the more efficiently the residual can be compressed
	 * by the DCT transforms, at the expense of more computation */
	MaxTUSize ffcommon.FUint32T

	/* The additional depth the residual quad-tree is allowed to recurse beyond
	 * the coding quad-tree, for inter coded blocks. This must be between 1 and
	 * 4. The higher the value the more efficiently the residual can be
	 * compressed by the DCT transforms, at the expense of much more compute */
	TuQTMaxInterDepth ffcommon.FUint32T

	/* The additional depth the residual quad-tree is allowed to recurse beyond
	 * the coding quad-tree, for intra coded blocks. This must be between 1 and
	 * 4. The higher the value the more efficiently the residual can be
	 * compressed by the DCT transforms, at the expense of much more compute */
	TuQTMaxIntraDepth ffcommon.FUint32T

	/* Enable early exit decisions for inter coded blocks to avoid recursing to
	 * higher TU depths. Default: 0 */
	LimitTU ffcommon.FUint32T

	/* Set the amount of rate-distortion analysis to use within quant. 0 implies
	 * no rate-distortion optimization. At level 1 rate-distortion cost is used to
	 * find optimal rounding values for each level (and allows psy-rdoq to be
	 * enabled). At level 2 rate-distortion cost is used to make decimate decisions
	 * on each 4x4 coding group (including the cost of signaling the group within
	 * the group bitmap).  Psy-rdoq is less effective at preserving energy when
	 * RDOQ is at level 2. Default: 0 */
	RdoqLevel ffcommon.FInt

	/* Enable the implicit signaling of the sign bit of the last coefficient of
	 * each transform unit. This saves one bit per TU at the expense of figuring
	 * out which coefficient can be toggled with the least distortion.
	 * Default is enabled */
	BEnableSignHiding ffcommon.FInt

	/* Allow intra coded blocks to be encoded directly as residual without the
	 * DCT transform, when this improves efficiency. Checking whether the block
	 * will benefit from this option incurs a performance penalty. Default is
	 * disabled */
	BEnableTransformSkip ffcommon.FInt

	/* An integer value in range of 0 to 2000, which denotes strength of noise
	 * reduction in intra CUs. 0 means disabled */
	NoiseReductionIntra ffcommon.FInt

	/* An integer value in range of 0 to 2000, which denotes strength of noise
	 * reduction in inter CUs. 0 means disabled */
	NoiseReductionInter ffcommon.FInt

	/* Quantization scaling lists. HEVC supports 6 quantization scaling lists to
	 * be defined; one each for Y, Cb, Cr for intra prediction and one each for
	 * inter prediction.
	 *
	 * - NULL and "off" will disable quant scaling (default)
	 * - "default" will enable the HEVC default scaling lists, which
	 *   do not need to be signaled since they are specified
	 * - all other strings indicate a filename containing custom scaling lists
	 *   in the HM format. The encode will fail if the file is not parsed
	 *   correctly. Custom lists must be signaled in the SPS. */
	ScalingLists ffcommon.FCharPStruct

	/*== Intra Coding Tools ==*/

	/* Enable constrained intra prediction. This causes intra prediction to
	 * input samples that were inter predicted. For some use cases this is
	 * believed to me more robust to stream errors, but it has a compression
	 * penalty on P and (particularly) B slices. Defaults to disabled */
	BEnableConstrainedIntra ffcommon.FInt

	/* Enable strong intra smoothing for 32x32 blocks where the reference
	 * samples are flat. It may or may not improve compression efficiency,
	 * depending on your source material. Defaults to disabled */
	BEnableStrongIntraSmoothing ffcommon.FInt

	/*== Inter Coding Tools ==*/

	/* The maximum number of merge candidates that are considered during inter
	 * analysis.  This number (between 1 and 5) is signaled in the stream
	 * headers and determines the number of bits required to signal a merge so
	 * it can have significant trade-offs. The smaller this number the higher
	 * the performance but the less compression efficiency. Default is 3 */
	MaxNumMergeCand ffcommon.FUint32T

	/* Limit the motion references used for each search based on the results of
	 * previous motion searches already performed for the same CU: If 0 all
	 * references are always searched. If X265_REF_LIMIT_CU all motion searches
	 * will restrict themselves to the references selected by the 2Nx2N search
	 * at the same depth. If X265_REF_LIMIT_DEPTH the 2Nx2N motion search will
	 * only use references that were selected by the best motion searches of the
	 * 4 split CUs at the next lower CU depth.  The two flags may be combined */
	LimitReferences ffcommon.FUint32T

	/* Limit modes analyzed for each CU using cost metrics from the 4 sub-CUs */
	LimitModes ffcommon.FUint32T

	/* ME search method (DIA, HEX, UMH, STAR, SEA, FULL). The search patterns
	 * (methods) are sorted in increasing complexity, with diamond being the
	 * simplest and fastest and full being the slowest.  DIA, HEX, UMH and SEA were
	 * adapted from x264 directly. STAR is an adaption of the HEVC reference
	 * encoder's three step search, while full is a naive exhaustive search. The
	 * default is the star search, it has a good balance of performance and
	 * compression efficiency */
	SearchMethod ffcommon.FInt

	/* A value between 0 and X265_MAX_SUBPEL_LEVEL which adjusts the amount of
	 * effort performed during sub-pel refine. Default is 5 */
	SubpelRefine ffcommon.FInt

	/* The maximum distance from the motion prediction that the full pel motion
	 * search is allowed to progress before terminating. This value can have an
	 * effect on frame parallelism, as referenced frames must be at least this
	 * many rows of reconstructed pixels ahead of the referencee at all times.
	 * (When considering reference lag, the motion prediction must be ignored
	 * because it cannot be known ahead of time).  Default is 60, which is the
	 * default max CU size (64) minus the luma HPEL half-filter length (4). If a
	 * smaller CU size is used, the search range should be similarly reduced */
	SearchRange ffcommon.FInt

	/* Enable availability of temporal motion vector for AMVP, default is enabled */
	BEnableTemporalMvp ffcommon.FInt

	/* Enable 3-level Hierarchical motion estimation at One-Sixteenth, Quarter and Full resolution.
	 * Default is disabled */
	BEnableHME ffcommon.FInt

	/* Enable HME search method (DIA, HEX, UMH, STAR, SEA, FULL) for level 0, 1 and 2.
	 * Default is hex, umh, umh for L0, L1 and L2 respectively. */
	HmeSearchMethod [3]ffcommon.FInt

	/* Enable weighted prediction in P slices.  This enables weighting analysis
	 * in the lookahead, which influences slice decisions, and enables weighting
	 * analysis in the main encoder which allows P reference samples to have a
	 * weight function applied to them prior to using them for motion
	 * compensation.  In video which has lighting changes, it can give a large
	 * improvement in compression efficiency. Default is enabled */
	BEnableWeightedPred ffcommon.FInt

	/* Enable weighted prediction in B slices. Default is disabled */
	BEnableWeightedBiPred ffcommon.FInt
	/* Enable source pixels in motion estimation. Default is disabled */
	BSourceReferenceEstimation ffcommon.FInt
	/*== Loop Filters ==*/
	/* Enable the deblocking loop filter, which improves visual quality by
	 * reducing blocking effects at block edges, particularly at lower bitrates
	 * or higher QP. When enabled it adds another CU row of reference lag,
	 * reducing frame parallelism effectiveness. Default is enabled */
	BEnableLoopFilter ffcommon.FInt

	/* deblocking filter tC offset [-6, 6] -6 light filter, 6 strong.
	 * This is the coded div2 value, actual offset is doubled at use */
	DeblockingFilterTCOffset ffcommon.FInt

	/* deblocking filter Beta offset [-6, 6] -6 light filter, 6 strong
	 * This is the coded div2 value, actual offset is doubled at use */
	DeblockingFilterBetaOffset ffcommon.FInt

	/* Enable the Sample Adaptive Offset loop filter, which reduces distortion
	 * effects by adjusting reconstructed sample values based on histogram
	 * analysis to better approximate the original samples. When enabled it adds
	 * a CU row of reference lag, reducing frame parallelism effectiveness.
	 * Default is enabled */
	BEnableSAO ffcommon.FInt

	/* Note: when deblocking and SAO are both enabled, the loop filter CU lag is
	 * only one row, as they operate in series on the same row. */

	/* Select the method in which SAO deals with deblocking boundary pixels. If
	 * disabled the right and bottom boundary areas are skipped. If enabled,
	 * non-deblocked pixels are used entirely. Default is disabled */
	BSaoNonDeblocked ffcommon.FInt

	/* Select tune rate in which SAO has to be applied.
	   1 - Filtering applied only on I-frames(I) [Light tune]
	   2 - No Filtering on B frames (I, P) [Medium tune]
	   3 - No Filtering on non-ref b frames  (I, P, B) [Strong tune] */
	SelectiveSAO ffcommon.FInt

	/*== Analysis tools ==*/

	/* A value between 1 and 6 (both inclusive) which determines the level of
	 * rate distortion optimizations to perform during mode and depth decisions.
	 * The more RDO the better the compression efficiency at a major cost of
	 * performance. Default is 3 */
	RdLevel ffcommon.FInt

	/* Enable early skip decisions to avoid analysing additional modes in likely
	 * skip blocks. Default is disabled */
	BEnableEarlySkip ffcommon.FInt

	/* Enable early CU size decisions to avoid recursing to higher depths.
	 * Default is enabled */
	RecursionSkipMode ffcommon.FInt

	/* Use a faster search method to find the best intra mode. Default is 0 */
	BEnableFastIntra ffcommon.FInt

	/* Enable a faster determination of whether skipping the DCT transform will
	 * be beneficial. Slight performance gain for some compression loss. Default
	 * is enabled */
	BEnableTSkipFast ffcommon.FInt

	/* The CU Lossless flag, when enabled, compares the rate-distortion costs
	 * for normal and lossless encoding, and chooses the best mode for each CU.
	 * If lossless mode is chosen, the cu-transquant-bypass flag is set for that
	 * CU */
	BCULossless ffcommon.FInt

	/* Specify whether to attempt to encode intra modes in B frames. By default
	 * enabled, but only applicable for the presets which use rdLevel 5 or 6
	 * (veryslow and placebo). All other presets will not try intra in B frames
	 * regardless of this setting */
	BIntraInBFrames ffcommon.FInt

	/* Apply an optional penalty to the estimated cost of 32x32 intra blocks in
	 * non-intra slices. 0 is disabled, 1 enables a small penalty, and 2 enables
	 * a full penalty. This favors inter-coding and its low bitrate over
	 * potential increases in distortion, but usually improves performance.
	 * Default is 0 */
	RdPenalty ffcommon.FInt

	/* Psycho-visual rate-distortion strength. Only has an effect in presets
	 * which use RDO. It makes mode decision favor options which preserve the
	 * energy of the source, at the cost of lost compression. The value must
	 * be between 0 and 5.0, 1.0 is typical. Default 2.0 */
	PsyRd ffcommon.FDouble

	/* Strength of psycho-visual optimizations in quantization. Only has an
	 * effect when RDOQ is enabled (presets slow, slower and veryslow). The
	 * value must be between 0 and 50, 1.0 is typical. Default 0 */
	PsyRdoq ffcommon.FDouble

	/* Perform quantisation parameter based RD refinement. RD cost is calculated
	 * on the best CU partitions, chosen after the CU analysis, for a range of QPs
	 * to find the optimal rounding effect. Only effective at rd-levels 5 and 6.
	 * Default disabled */
	BEnableRdRefine ffcommon.FInt

	/* If save, write per-frame analysis information into analysis buffers.
	 * If load, read analysis information into analysis buffer and use this
	 * analysis information to reduce the amount of work the encoder must perform.
	 * Default disabled. Now deprecated*/
	AnalysisReuseMode ffcommon.FInt

	/* Filename for multi-pass-opt-analysis/distortion. Default name is "x265_analysis.dat" */
	AnalysisReuseFileName ffcommon.FCharPStruct

	/*== Rate Control ==*/

	/* The lossless flag enables true lossless coding, bypassing scaling,
	 * transform, quantization and in-loop filter processes. This is used for
	 * ultra-high bitrates with zero loss of quality. It implies no rate control */
	BLossless ffcommon.FInt

	/* Generally a small signed integer which offsets the QP used to quantize
	 * the Cb chroma residual (delta from luma QP specified by rate-control).
	 * Default is 0, which is recommended */
	CbQpOffset ffcommon.FInt

	/* Generally a small signed integer which offsets the QP used to quantize
	 * the Cr chroma residual (delta from luma QP specified by rate-control).
	 * Default is 0, which is recommended */
	CrQpOffset ffcommon.FInt

	/* Specifies the preferred transfer characteristics syntax element in the
	 * alternative transfer characteristics SEI message (see. D.2.38 and D.3.38 of
	 * JCTVC-W1005 http://phenix.it-sudparis.eu/jct/doc_end_user/documents/23_San%20Diego/wg11/JCTVC-W1005-v4.zip
	 * */
	PreferredTransferCharacteristics ffcommon.FInt

	/*
	 * Specifies the value for the pic_struc syntax element of the picture timing SEI message (See D2.3 and D3.3)
	 * of the HEVC spec. for a detailed explanation
	 * */
	PictureStructure ffcommon.FInt

	Rc struct {
		/* Explicit mode of rate-control, necessary for API users. It must
		 * be one of the X265_RC_METHODS enum values. */
		RateControlMode ffcommon.FInt

		/* Base QP to use for Constant QP rate control. Adaptive QP may alter
		 * the QP used for each block. If a QP is specified on the command line
		 * CQP rate control is implied. Default: 32 */
		Qp ffcommon.FInt

		/* target bitrate for Average BitRate (ABR) rate control. If a non- zero
		 * bitrate is specified on the command line, ABR is implied. Default 0 */
		Bitrate ffcommon.FInt

		/* qComp sets the quantizer curve compression factor. It weights the frame
		 * quantizer based on the complexity of residual (measured by lookahead).
		 * Default value is 0.6. Increasing it to 1 will effectively generate CQP */
		QCompress ffcommon.FDouble

		/* QP offset between I/P and P/B frames. Default ipfactor: 1.4
		 * Default pbFactor: 1.3 */
		IpFactor ffcommon.FDouble
		PbFactor ffcommon.FDouble

		/* Ratefactor constant: targets a certain constant "quality".
		 * Acceptable values between 0 and 51. Default value: 28 */
		RfConstant ffcommon.FDouble

		/* Max QP difference between frames. Default: 4 */
		QpStep ffcommon.FInt

		/* Enable adaptive quantization. This mode distributes available bits between all
		 * CTUs of a frame, assigning more bits to low complexity areas. Turning
		 * this ON will usually affect PSNR negatively, however SSIM and visual quality
		 * generally improves. Default: X265_AQ_AUTO_VARIANCE */
		AqMode ffcommon.FInt

		/*
		 * Enable adaptive quantization.
		 * It scales the quantization step size according to the spatial activity of one
		 * coding unit relative to frame average spatial activity. This AQ method utilizes
		 * the minimum variance of sub-unit in each coding unit to represent the coding
		 * unit’s spatial complexity. */
		HevcAq ffcommon.FInt

		/* Sets the strength of AQ bias towards low detail CTUs. Valid only if
		 * AQ is enabled. Default value: 1.0. Acceptable values between 0.0 and 3.0 */
		AqStrength ffcommon.FDouble

		/* Delta QP range by QP adaptation based on a psycho-visual model.
		 * Acceptable values between 1.0 to 6.0 */
		QpAdaptationRange ffcommon.FDouble

		/* Sets the maximum rate the VBV buffer should be assumed to refill at
		 * Default is zero */
		VbvMaxBitrate ffcommon.FInt

		/* Sets the size of the VBV buffer in kilobits. Default is zero */
		VbvBufferSize ffcommon.FInt

		/* Sets how full the VBV buffer must be before playback starts. If it is less than
		 * 1, then the initial fill is vbv-init * vbvBufferSize. Otherwise, it is
		 * interpreted as the initial fill in kbits. Default is 0.9 */
		VbvBufferInit ffcommon.FDouble

		/* Enable CUTree rate-control. This keeps track of the CUs that propagate temporally
		 * across frames and assigns more bits to these CUs. Improves encode efficiency.
		 * Default: enabled */
		CuTree ffcommon.FInt

		/* In CRF mode, maximum CRF as caused by VBV. 0 implies no limit */
		RfConstantMax ffcommon.FDouble

		/* In CRF mode, minimum CRF as caused by VBV */
		RfConstantMin ffcommon.FDouble

		/* Multi-pass encoding */
		/* Enable writing the stats in a multi-pass encode to the stat output file/memory */
		BStatWrite ffcommon.FInt

		/* Enable loading data from the stat input file/memory in a multi pass encode */
		BStatRead ffcommon.FInt

		/* Filename of the 2pass output/input stats file, if unspecified the
		 * encoder will default to using x265_2pass.log */
		StatFileName ffcommon.FCharPStruct

		/* temporally blur quants */
		Qblur ffcommon.FDouble

		/* temporally blur complexity */
		ComplexityBlur ffcommon.FDouble

		/* Enable slow and a more detailed first pass encode in multi pass rate control */
		BEnableSlowFirstPass ffcommon.FInt

		/* rate-control overrides */
		ZoneCount ffcommon.FInt
		Zones     *X265Zone

		/* number of zones in zone-file*/
		ZonefileCount ffcommon.FInt

		/* specify a text file which contains MAX_MAX_QP + 1 floating point
		 * values to be copied into x265_lambda_tab and a second set of
		 * MAX_MAX_QP + 1 floating point values for x265_lambda2_tab. All values
		 * are separated by comma, space or newline. Text after a hash (#) is
		 * ignored. The lambda tables are process-global, so these new lambda
		 * values will affect all encoders in the same process */
		LambdaFileName ffcommon.FCharPStruct

		/* Enable stricter conditions to check bitrate deviations in CBR mode. May compromise
		 * quality to maintain bitrate adherence */
		BStrictCbr ffcommon.FInt

		/* Enable adaptive quantization at CU granularity. This parameter specifies
		 * the minimum CU size at which QP can be adjusted, i.e. Quantization Group
		 * (QG) size. Allowed values are 64, 32, 16, 8 provided it falls within the
		 * inclusuve range [maxCUSize, minCUSize]. Experimental, default: maxCUSize */
		QgSize ffcommon.FUint32T

		/* internally enable if tune grain is set */
		BEnableGrain ffcommon.FInt

		/* sets a hard upper limit on QP */
		QpMax ffcommon.FInt

		/* sets a hard lower limit on QP */
		QpMin ffcommon.FInt

		/* internally enable if tune grain is set */
		BEnableConstVbv ffcommon.FInt

		/* if only the focused frames would be re-encode or not */
		BEncFocusedFramesOnly ffcommon.FInt

		/* Share the data with stats file or shared memory.
		   It must be one of the X265_DATA_SHARE_MODES enum values
		   Available if the bStatWrite or bStatRead is true.
		   Use stats file by default.
		   The stats file mode would be used among the encoders running in sequence.
		   The shared memory mode could only be used among the encoders running in parallel.
		   Now only the cutree data could be shared among shared memory. More data would be support in the future.*/
		DataShareMode ffcommon.FInt

		/* Unique shared memory name. Required if the shared memory mode enabled. NULL by default */
		SharedMemName ffcommon.FCharPStruct
	}

	/*== Video Usability Information ==*/
	Vui struct {
		/* Aspect ratio idc to be added to the VUI.  The default is 0 indicating
		 * the apsect ratio is unspecified. If set to X265_EXTENDED_SAR then
		 * sarWidth and sarHeight must also be set */
		AspectRatioIdc ffcommon.FInt

		/* Sample Aspect Ratio width in arbitrary units to be added to the VUI
		 * only if aspectRatioIdc is set to X265_EXTENDED_SAR.  This is the width
		 * of an individual pixel. If this is set then sarHeight must also be set */
		SarWidth ffcommon.FInt

		/* Sample Aspect Ratio height in arbitrary units to be added to the VUI.
		 * only if aspectRatioIdc is set to X265_EXTENDED_SAR.  This is the width
		 * of an individual pixel. If this is set then sarWidth must also be set */
		SarHeight ffcommon.FInt

		/* Enable overscan info present flag in the VUI.  If this is set then
		 * bEnabledOverscanAppropriateFlag will be added to the VUI. The default
		 * is false */
		BEnableOverscanInfoPresentFlag ffcommon.FInt

		/* Enable overscan appropriate flag.  The status of this flag is added
		 * to the VUI only if bEnableOverscanInfoPresentFlag is set. If this
		 * flag is set then cropped decoded pictures may be output for display.
		 * The default is false */
		BEnableOverscanAppropriateFlag ffcommon.FInt

		/* Video signal type present flag of the VUI.  If this is set then
		 * videoFormat, bEnableVideoFullRangeFlag and
		 * bEnableColorDescriptionPresentFlag will be added to the VUI. The
		 * default is false */
		BEnableVideoSignalTypePresentFlag ffcommon.FInt

		/* Video format of the source video.  0 = component, 1 = PAL, 2 = NTSC,
		 * 3 = SECAM, 4 = MAC, 5 = unspecified video format is the default */
		VideoFormat ffcommon.FInt

		/* Video full range flag indicates the black level and range of the luma
		 * and chroma signals as derived from E′Y, E′PB, and E′PR or E′R, E′G,
		 * and E′B real-valued component signals. The default is false */
		BEnableVideoFullRangeFlag ffcommon.FInt

		/* Color description present flag in the VUI. If this is set then
		 * color_primaries, transfer_characteristics and matrix_coeffs are to be
		 * added to the VUI. The default is false */
		BEnableColorDescriptionPresentFlag ffcommon.FInt

		/* Color primaries holds the chromacity coordinates of the source
		 * primaries. The default is 2 */
		ColorPrimaries ffcommon.FInt

		/* Transfer characteristics indicates the opto-electronic transfer
		 * characteristic of the source picture. The default is 2 */
		TransferCharacteristics ffcommon.FInt

		/* Matrix coefficients used to derive the luma and chroma signals from
		 * the red, blue and green primaries. The default is 2 */
		MatrixCoeffs ffcommon.FInt

		/* Chroma location info present flag adds chroma_sample_loc_type_top_field and
		 * chroma_sample_loc_type_bottom_field to the VUI. The default is false */
		BEnableChromaLocInfoPresentFlag ffcommon.FInt

		/* Chroma sample location type top field holds the chroma location in
		 * the top field. The default is 0 */
		ChromaSampleLocTypeTopField ffcommon.FInt

		/* Chroma sample location type bottom field holds the chroma location in
		 * the bottom field. The default is 0 */
		ChromaSampleLocTypeBottomField ffcommon.FInt

		/* Default display window flag adds def_disp_win_left_offset,
		 * def_disp_win_right_offset, def_disp_win_top_offset and
		 * def_disp_win_bottom_offset to the VUI. The default is false */
		BEnableDefaultDisplayWindowFlag ffcommon.FInt

		/* Default display window left offset holds the left offset with the
		 * conformance cropping window to further crop the displayed window */
		DefDispWinLeftOffset ffcommon.FInt

		/* Default display window right offset holds the right offset with the
		 * conformance cropping window to further crop the displayed window */
		DefDispWinRightOffset ffcommon.FInt

		/* Default display window top offset holds the top offset with the
		 * conformance cropping window to further crop the displayed window */
		DefDispWinTopOffset ffcommon.FInt

		/* Default display window bottom offset holds the bottom offset with the
		 * conformance cropping window to further crop the displayed window */
		DefDispWinBottomOffset ffcommon.FInt
	}

	/* SMPTE ST 2086 mastering display color volume SEI info, specified as a
	 * string which is parsed when the stream header SEI are emitted. The string
	 * format is "G(%hu,%hu)B(%hu,%hu)R(%hu,%hu)WP(%hu,%hu)L(%u,%u)" where %hu
	 * are unsigned 16bit integers and %u are unsigned 32bit integers. The SEI
	 * includes X,Y display primaries for RGB channels, white point X,Y and
	 * max,min luminance values. */
	MasteringDisplayColorVolume ffcommon.FCharPStruct

	/* Maximum Content light level(MaxCLL), specified as integer that indicates the
	 * maximum pixel intensity level in units of 1 candela per square metre of the
	 * bitstream. x265 will also calculate MaxCLL programmatically from the input
	 * pixel values and set in the Content light level info SEI */
	MaxCLL ffcommon.FUint16T

	/* Maximum Frame Average Light Level(MaxFALL), specified as integer that indicates
	 * the maximum frame average intensity level in units of 1 candela per square
	 * metre of the bitstream. x265 will also calculate MaxFALL programmatically
	 * from the input pixel values and set in the Content light level info SEI */
	MaxFALL ffcommon.FUint16T

	/* Minimum luma level of input source picture, specified as a integer which
	 * would automatically increase any luma values below the specified --min-luma
	 * value to that value. */
	MinLuma ffcommon.FUint16T

	/* Maximum luma level of input source picture, specified as a integer which
	 * would automatically decrease any luma values above the specified --max-luma
	 * value to that value. */
	MaxLuma ffcommon.FUint16T

	/* Maximum of the picture order count */
	Log2MaxPocLsb ffcommon.FInt

	/* Emit VUI Timing info, an optional VUI field */
	BEmitVUITimingInfo ffcommon.FInt

	/* Emit HRD Timing info */
	BEmitVUIHRDInfo ffcommon.FInt

	/* Maximum count of Slices of picture, the value range is [1, maximum rows] */
	MaxSlices ffcommon.FUnsignedInt

	/* Optimize QP in PPS based on statistics from prevvious GOP*/
	BOptQpPPS ffcommon.FInt

	/* Opitmize ref list length in PPS based on stats from previous GOP*/
	BOptRefListLengthPPS ffcommon.FInt

	/* Enable storing commonly RPS in SPS in multi pass mode */
	BMultiPassOptRPS ffcommon.FInt
	/* This value represents the percentage difference between the inter cost and
	 * intra cost of a frame used in scenecut detection. Default 5. */
	ScenecutBias ffcommon.FDouble
	/* Use multiple worker threads dedicated to doing only lookahead instead of sharing
	 * the worker threads with Frame Encoders. A dedicated lookahead threadpool is created with the
	 * specified number of worker threads. This can range from 0 upto half the
	 * hardware threads available for encoding. Using too many threads for lookahead can starve
	 * resources for frame Encoder and can harm performance. Default is 0 - disabled. */
	LookaheadThreads ffcommon.FInt
	/* Optimize CU level QPs to signal consistent deltaQPs in frame for rd level > 4 */
	BOptCUDeltaQP ffcommon.FInt
	/* Refine analysis in multipass ratecontrol based on analysis information stored */
	AnalysisMultiPassRefine ffcommon.FInt
	/* Refine analysis in multipass ratecontrol based on distortion data stored */
	AnalysisMultiPassDistortion ffcommon.FInt
	/* Adaptive Quantization based on relative motion */
	BAQMotion ffcommon.FInt
	/* SSIM based RDO, based on residual divisive normalization scheme. Used for mode
	 * selection during analysis of CTUs, can achieve significant gain in terms of
	 * objective quality metrics SSIM and PSNR */
	BSsimRd ffcommon.FInt

	/* Increase RD at points where bitrate drops due to vbv. Default 0 */
	DynamicRd ffcommon.FDouble

	/* Enables the emitting of HDR SEI packets which contains HDR-specific params.
	 * Auto-enabled when max-cll, max-fall, or mastering display info is specified.
	 * Default is disabled. Now deprecated.*/
	BEmitHDRSEI ffcommon.FInt

	/* Enable luma and chroma offsets for HDR/WCG content.
	 * Default is disabled. Now deprecated.*/
	BHDROpt ffcommon.FInt

	/* A value between 1 and 10 (both inclusive) determines the level of
	 * information stored/reused in analysis save/load. Higher the refine
	 * level higher the information stored/reused. Default is 5. Now deprecated. */
	AnalysisReuseLevel ffcommon.FInt

	/* Limit Sample Adaptive Offset filter computation by early terminating SAO
	 * process based on inter prediction mode, CTU spatial-domain correlations,
	 * and relations between luma and chroma */
	BLimitSAO ffcommon.FInt

	/* File containing the tone mapping information */
	ToneMapFile ffcommon.FCharPStruct

	/* Insert tone mapping information only for IDR frames and when the
	 * tone mapping information changes. */
	BDhdr10opt ffcommon.FInt

	/* Determine how x265 react to the content information recieved through the API */
	BCTUInfo ffcommon.FInt

	/* Use ratecontrol statistics from pic_in, if available*/
	BUseRcStats ffcommon.FInt

	/* Factor by which input video is scaled down for analysis save mode. Default is 0 */
	ScaleFactor ffcommon.FInt

	/* Enable intra refinement in load mode*/
	IntraRefine ffcommon.FInt

	/* Enable inter refinement in load mode*/
	InterRefine ffcommon.FInt

	/* Enable motion vector refinement in load mode*/
	MvRefine ffcommon.FInt

	/* Log of maximum CTU size */
	MaxLog2CUSize ffcommon.FUint32T

	/* Actual CU depth with respect to config depth */
	MaxCUDepth ffcommon.FUint32T

	/* CU depth with respect to maximum transform size */
	UnitSizeDepth ffcommon.FUint32T

	/* Number of 4x4 units in maximum CU size */
	Num4x4Partitions ffcommon.FUint32T

	/* Specify if analysis mode uses file for data reuse */
	BUseAnalysisFile ffcommon.FInt

	/* File pointer for csv log */
	Csvfpt *FILE

	/* Force flushing the frames from encoder */
	ForceFlush ffcommon.FInt

	/* Enable skipping split RD analysis when sum of split CU rdCost larger than none split CU rdCost for Intra CU */
	BEnableSplitRdSkip ffcommon.FInt

	/* Disable lookahead */
	BDisableLookahead ffcommon.FInt

	/* Use low-pass subband dct approximation
	 *  This DCT approximation is less computational intensive and gives results close to standard DCT */
	BLowPassDct ffcommon.FInt

	/* Sets the portion of the decode buffer that must be available after all the
	 * specified frames have been inserted into the decode buffer. If it is less
	 * than 1, then the final buffer available is vbv-end * vbvBufferSize.  Otherwise,
	 * it is interpreted as the final buffer available in kbits. Default 0 (disabled) */
	VbvBufferEnd ffcommon.FDouble

	/* Frame from which qp has to be adjusted to hit final decode buffer emptiness.
	 * Specified as a fraction of the total frames. Default 0 */
	VbvEndFrameAdjust ffcommon.FDouble

	/* Reuse MV information obtained through API */
	BAnalysisType ffcommon.FInt
	/* Allow the encoder to have a copy of the planes of x265_picture in Frame */
	BCopyPicToFrame ffcommon.FInt

	/*Number of frames for GOP boundary decision lookahead.If a scenecut frame is found
	 * within this from the gop boundary set by keyint, the GOP will be extented until such a point,
	 * otherwise the GOP will be terminated as set by keyint*/
	GopLookahead ffcommon.FInt

	/*Write per-frame analysis information into analysis buffers. Default disabled. */
	AnalysisSave ffcommon.FCharPStruct

	/* Read analysis information into analysis buffer and use this analysis information
	 * to reduce the amount of work the encoder must perform. Default disabled. */
	AnalysisLoad ffcommon.FCharPStruct

	/*Number of RADL pictures allowed in front of IDR*/
	Radl ffcommon.FInt

	/* This value controls the maximum AU size defined in specification
	 * It represents the percentage of maximum AU size used.
	 * Default is 1 (which is 100%). Range is 0.5 to 1. */
	MaxAUSizeFactor ffcommon.FDouble

	/* Enables the emission of a Recovery Point SEI with the stream headers
	 * at each IDR frame describing poc of the recovery point, exact matching flag
	 * and broken link flag. Default is disabled. */
	BEmitIDRRecoverySEI ffcommon.FInt

	/* Dynamically change refine-inter at block level*/
	BDynamicRefine ffcommon.FInt

	/* Enable writing all SEI messgaes in one single NAL instead of mul*/
	BSingleSeiNal ffcommon.FInt

	/* First frame of the chunk. Frames preceeding this in display order will
	 * be encoded, however, they will be discarded in the bitstream.
	 * Default 0 (disabled). */
	ChunkStart ffcommon.FInt

	/* Last frame of the chunk. Frames following this in display order will be
	 * used in taking lookahead decisions, but, they will not be encoded.
	 * Default 0 (disabled). */
	ChunkEnd ffcommon.FInt
	/* File containing base64 encoded SEI messages in POC order */
	NaluFile ffcommon.FCharPStruct

	/* Generate bitstreams confirming to the specified dolby vision profile,
	 * note that 0x7C01 makes RPU appear to be an unspecified NAL type in
	 * HEVC stream. if BL is backward compatible, Dolby Vision single
	 * layer VES will be equivalent to a backward compatible BL VES on legacy
	 * device as RPU will be ignored. Default 0 (disabled) */
	DolbyProfile ffcommon.FInt

	/* Set concantenation flag for the first keyframe in the HRD buffering period SEI. */
	BEnableHRDConcatFlag ffcommon.FInt

	/* Store/normalize ctu distortion in analysis-save/load. Ranges from 0 - 1.
	 *  0 - Disabled. 1 - Save/Load ctu distortion to/from the file specified
	 * analysis-save/load. Default 0. */
	CtuDistortionRefine ffcommon.FInt

	/* Enable SVT HEVC Encoder */
	BEnableSvtHevc ffcommon.FInt

	/* SVT-HEVC param structure. For internal use when SVT HEVC encoder is enabled */
	SvtHevcParam ffcommon.FVoidP

	/* Detect fade-in regions. Enforces I-slice for the brightest point.
	   Re-init RC history at that point in ABR mode. Default is disabled. */
	BEnableFades ffcommon.FInt

	/* Enable field coding */
	BField ffcommon.FInt

	/*Emit content light level info SEI*/
	BEmitCLL ffcommon.FInt

	/*
	 * Signals picture structure SEI timing message for every frame
	 * picture structure 7 is signalled for frame doubling
	 * picture structure 8 is signalled for frame tripling
	 * */
	BEnableFrameDuplication ffcommon.FInt

	/*
	 * For adaptive frame duplication, a threshold is set above which the frames are similar.
	 * User can set a variable threshold. Default 70.
	 * */
	DupThreshold ffcommon.FInt

	/*Input sequence bit depth. It can be either 8bit, 10bit or 12bit.*/
	SourceBitDepth ffcommon.FInt

	/*Size of the zone to be reconfigured in frames. Default 0. API only. */
	ReconfigWindowSize ffcommon.FUint32T

	/*Flag to indicate if rate-control history has to be reset during zone reconfiguration.
	  Default 1 (Enabled). API only. */
	BResetZoneConfig ffcommon.FInt

	/* It reduces the bits spent on the inter-frames within the scenecutWindow before and / or after a scenecut
	 * by increasing their QP in ratecontrol pass2 algorithm without any deterioration in visual quality.
	 * 0 - Disabled (default).
	 * 1 - Forward masking.
	 * 2 - Backward masking.
	 * 3 - Bi-directional masking. */
	BEnableSceneCutAwareQp ffcommon.FInt

	/* The duration(in milliseconds) for which there is a reduction in the bits spent on the inter-frames after a scenecut
	 * by increasing their QP, when bEnableSceneCutAwareQp is 1 or 3. Default is 500ms.*/
	FwdMaxScenecutWindow ffcommon.FInt
	FwdScenecutWindow    [6]ffcommon.FInt

	/* The offset by which QP is incremented for inter-frames after a scenecut when bEnableSceneCutAwareQp is 1 or 3.
	 * Default is +5. */
	FwdRefQpDelta [6]ffcommon.FDouble

	/* The offset by which QP is incremented for non-referenced inter-frames after a scenecut when bEnableSceneCutAwareQp is 1 or 3. */
	FwdNonRefQpDelta [6]ffcommon.FDouble

	/* Enables histogram based scenecut detection algorithm to detect scenecuts. Default disabled */
	BHistBasedSceneCut ffcommon.FInt

	/* Enable HME search ranges for L0, L1 and L2 respectively. */
	HmeRange [3]ffcommon.FInt

	/* Block-level QP optimization for HDR10 content. Default is disabled.*/
	BHDR10Opt ffcommon.FInt

	/* Enables the emitting of HDR10 SEI packets which contains HDR10-specific params.
	 * Auto-enabled when max-cll, max-fall, or mastering display info is specified.
	 * Default is disabled */
	BEmitHDR10SEI ffcommon.FInt

	/* A value between 1 and 10 (both inclusive) determines the level of
	 * analysis information stored in analysis-save. Higher the refine level higher
	 * the information stored. Default is 5 */
	AnalysisSaveReuseLevel ffcommon.FInt

	/* A value between 1 and 10 (both inclusive) determines the level of
	 * analysis information reused in analysis-load. Higher the refine level higher
	 * the information reused. Default is 5 */
	AnalysisLoadReuseLevel ffcommon.FInt

	/* Conformance window right offset specifies the padding offset to the
	 * right side of the internal copy of the input pictures in the library.
	 * The decoded picture will be cropped based on conformance window right offset
	 * signaled in the SPS before output. Default is 0.
	 * Recommended to set this during non-file based analysis-load.
	 * This is to inform the encoder about the conformace window right offset
	 * to be added to match the number of CUs across the width for which analysis
	 * info is available from the corresponding analysis-save. */

	ConfWinRightOffset ffcommon.FInt

	/* Conformance window bottom offset specifies the padding offset to the
	 * bottom side of the internal copy of the input pictures in the library.
	 * The decoded picture will be cropped based on conformance window bottom offset
	 * signaled in the SPS before output. Default is 0.
	 * Recommended to set this during non-file based analysis-load.
	 * This is to inform the encoder about the conformace window bottom offset
	 * to be added to match the number of CUs across the height for which analysis
	 * info is available from the corresponding analysis-save. */

	ConfWinBottomOffset ffcommon.FInt

	/* Edge variance threshold for quad tree establishment. */
	EdgeVarThreshold ffcommon.FFloat

	/* Maxrate that could be signaled to the decoder. Default 0. API only. */
	DecoderVbvMaxRate ffcommon.FInt

	/*Enables Qp tuning with respect to real time VBV buffer fullness in rate
	  control 2 pass. Experimental.Default is disabled*/
	BliveVBV2pass ffcommon.FInt

	/* Minimum VBV fullness to be maintained. Default 50. Keep the buffer
	 * at least 50% full */
	MinVbvFullness ffcommon.FDouble

	/* Maximum VBV fullness to be maintained. Default 80. Keep the buffer
	 * at max 80% full */
	MaxVbvFullness ffcommon.FDouble

	/* The duration(in milliseconds) for which there is a reduction in the bits spent on the inter-frames before a scenecut
	 * by increasing their QP, when bEnableSceneCutAwareQp is 2 or 3. Default is 100ms.*/
	BwdMaxScenecutWindow ffcommon.FInt
	BwdScenecutWindow    [6]ffcommon.FInt

	/* The offset by which QP is incremented for inter-frames before a scenecut when bEnableSceneCutAwareQp is 2 or 3. */
	BwdRefQpDelta [6]ffcommon.FDouble

	/* The offset by which QP is incremented for non-referenced inter-frames before a scenecut when bEnableSceneCutAwareQp is 2 or 3. */
	BwdNonRefQpDelta [6]ffcommon.FDouble

	/* Specify combinations of color primaries, transfer characteristics, color matrix,
	 * range of luma and chroma signals, and chroma sample location. This has higher
	 * precedence than individual VUI parameters. If any individual VUI option is specified
	 * together with this, which changes the values set corresponding to the system-id
	 * or color-volume, it will be discarded. */
	VideoSignalTypePreset ffcommon.FCharPStruct

	/* Flag indicating whether the encoder should emit an End of Bitstream
	 * NAL at the end of bitstream. Default false */
	BEnableEndOfBitstream ffcommon.FInt

	/* Flag indicating whether the encoder should emit an End of Sequence
	 * NAL at the end of every Coded Video Sequence. Default false */
	BEnableEndOfSequence ffcommon.FInt

	/* Film Grain Characteristic file */
	FilmGrain ffcommon.FCharPStruct

	/*Motion compensated temporal filter*/
	BEnableTemporalFilter  ffcommon.FInt
	TemporalFilterStrength ffcommon.FDouble

	/*SBRC*/
	BEnableSBRC ffcommon.FInt
}

// /* x265_param_alloc:
//   - Allocates an x265_param instance. The returned param structure is not
//   - special in any way, but using this method together with x265_param_free()
//   - and x265_param_parse() to set values by name allows the application to treat
//   - x265_param as an opaque data struct for version safety */
//
// x265_param *x265_param_alloc(void);
func X265ParamAlloc() (res *X265Param) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_param_alloc").Call()
	res = (*X265Param)(unsafe.Pointer(t))
	return
}

// /* x265_param_free:
//   - Use x265_param_free() to release storage for an x265_param instance
//   - allocated by x265_param_alloc() */
//
// void x265_param_free(x265_param *);
func (this *X265Param) X265ParamFree() {
	libx265common.GetLibx265Dll().NewProc("x265_param_free").Call(
		uintptr(unsafe.Pointer(this)),
	)
}

// /* x265_param_default:
//   - Initialize an x265_param structure to default values */
//
// void x265_param_default(x265_param *param);
func (this *X265Param) X265ParamDefault() {
	libx265common.GetLibx265Dll().NewProc("x265_param_default").Call(
		uintptr(unsafe.Pointer(this)),
	)
}

/* x265_param_parse:
 *  set one parameter by name.
 *  returns 0 on success, or returns one of the following errors.
 *  note: BAD_VALUE occurs only if it can't even parse the value,
 *  numerical range is not checked until x265_encoder_open().
 *  value=NULL means "true" for boolean options, but is a BAD_VALUE for non-booleans. */
const X265_PARAM_BAD_NAME = (-1)
const X265_PARAM_BAD_VALUE = (-2)

// int x265_param_parse(x265_param *p, const char *name, const char *value);
func (p *X265Param) X265ParamParse(name, value ffcommon.FCharP) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_param_parse").Call(
		uintptr(unsafe.Pointer(p)),
		ffcommon.UintPtrFromString(name),
		ffcommon.UintPtrFromString(value),
	)
	res = ffcommon.FInt(t)
	return
}

// x265_zone *x265_zone_alloc(int zoneCount, int isZoneFile);
func X265ZoneAlloc(zoneCount, isZoneFile ffcommon.FInt) (res *X265Zone) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_zone_alloc").Call(
		uintptr(zoneCount),
		uintptr(isZoneFile),
	)
	res = (*X265Zone)(unsafe.Pointer(t))
	return
}

// void x265_zone_free(x265_param *param);
func (param *X265Param) X265ZoneFree() {
	libx265common.GetLibx265Dll().NewProc("x265_zone_free").Call(
		uintptr(unsafe.Pointer(param)),
	)
}

// int x265_zone_param_parse(x265_param* p, const char* name, const char* value);
func (p *X265Param) X265ZoneParamParse(name, value ffcommon.FCharP) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_zone_param_parse").Call(
		uintptr(unsafe.Pointer(p)),
		ffcommon.UintPtrFromString(name),
		ffcommon.UintPtrFromString(value),
	)
	res = ffcommon.FInt(t)
	return
}

// int x265_scenecut_aware_qp_param_parse(x265_param* p, const char* name, const char* value);
func (p *X265Param) X265ScenecutAwareQpParamParse(name, value ffcommon.FCharP) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_scenecut_aware_qp_param_parse").Call(
		uintptr(unsafe.Pointer(p)),
		ffcommon.UintPtrFromString(name),
		ffcommon.UintPtrFromString(value),
	)
	res = ffcommon.FInt(t)
	return
}

var X265ProfileNames = []string{
	/* HEVC v1 */
	"main", "main10", "mainstillpicture" /* alias */, "msp",

	/* HEVC v2 (Range Extensions) */
	"main-intra", "main10-intra",
	"main444-8", "main444-intra", "main444-stillpicture",

	"main422-10", "main422-10-intra",
	"main444-10", "main444-10-intra",

	"main12", "main12-intra",
	"main422-12", "main422-12-intra",
	"main444-12", "main444-12-intra",

	"main444-16-intra", "main444-16-stillpicture", /* Not Supported! */

}

// /* x265_param_apply_profile:
//   - Applies the restrictions of the given profile. (one of x265_profile_names)
//   - (can be NULL, in which case the function will do nothing)
//   - Note: the detected profile can be lower than the one specified to this
//   - function. This function will force the encoder parameters to fit within
//   - the specified profile, or fail if that is impossible.
//   - returns 0 on success, negative on failure (e.g. invalid profile name). */
//
// int x265_param_apply_profile(x265_param *, const char *profile);
func (this *X265Param) X265ParamApplyProfile(profile ffcommon.FCharP) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_param_apply_profile").Call(
		uintptr(unsafe.Pointer(this)),
		ffcommon.UintPtrFromString(profile),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_param_default_preset:
//   - The same as x265_param_default, but also use the passed preset and tune
//   - to modify the default settings.
//   - (either can be NULL, which implies no preset or no tune, respectively)
//     *
//   - Currently available presets are, ordered from fastest to slowest: */
//
// static const char * const x265_preset_names[] = { "ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow", "placebo", 0 };
var X265PresetNames = []string{"ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow", "placebo"}

// /*      The presets can also be indexed numerically, as in:
//   - x265_param_default_preset( &param, "3", ... )
//   - with ultrafast mapping to "0" and placebo mapping to "9".  This mapping may
//   - of course change if new presets are added in between, but will always be
//   - ordered from fastest to slowest.
//     *
//   - Warning: the speed of these presets scales dramatically.  Ultrafast is a full
//   - 100 times faster than placebo!
//     *
//   - Currently available tunings are: */
//
// static const char * const x265_tune_names[] = { "psnr", "ssim", "grain", "zerolatency", "fastdecode", "animation", 0 };
var X265TuneNames = []string{"psnr", "ssim", "grain", "zerolatency", "fastdecode", "animation"}

// /*      returns 0 on success, negative on failure (e.g. invalid preset/tune name). */
// int x265_param_default_preset(x265_param *, const char *preset, const char *tune);

// /* x265_picture_alloc:
//   - Allocates an x265_picture instance. The returned picture structure is not
//   - special in any way, but using this method together with x265_picture_free()
//   - and x265_picture_init() allows some version safety. New picture fields will
//   - always be added to the end of x265_picture */
//
// x265_picture *x265_picture_alloc(void);
func X265PictureAlloc() (res *X265Picture) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_picture_alloc").Call()
	res = (*X265Picture)(unsafe.Pointer(t))
	return
}

// /* x265_picture_free:
//   - Use x265_picture_free() to release storage for an x265_picture instance
//   - allocated by x265_picture_alloc() */
//
// void x265_picture_free(x265_picture *);
func (this *X265Picture) X265PictureFree() {
	libx265common.GetLibx265Dll().NewProc("x265_picture_free").Call(
		uintptr(unsafe.Pointer(this)),
	)
}

// /* x265_picture_init:
//   - Initialize an x265_picture structure to default values. It sets the pixel
//   - depth and color space to the encoder's internal values and sets the slice
//   - type to auto - so the lookahead will determine slice type. */
//
// void x265_picture_init(x265_param *param, x265_picture *pic);
func (param *X265Param) X265PictureInit(pic *X265Picture) {
	libx265common.GetLibx265Dll().NewProc("x265_picture_init").Call(
		uintptr(unsafe.Pointer(param)),
		uintptr(unsafe.Pointer(pic)),
	)
}

// /* x265_max_bit_depth:
//  *      Specifies the numer of bits per pixel that x265 uses internally to
//  *      represent a pixel, and the bit depth of the output bitstream.
//  *      param->internalBitDepth must be set to this value. x265_max_bit_depth
//  *      will be 8 for default builds, 10 for HIGH_BIT_DEPTH builds. */
// X265_API extern const int x265_max_bit_depth;

// /* x265_version_str:
//  *      A static string containing the version of this compiled x265 library */
// X265_API extern const char *x265_version_str;

// /* x265_build_info:
//  *      A static string describing the compiler and target architecture */
// X265_API extern const char *x265_build_info_str;

/* x265_alloc_analysis_data:
*     Allocate memory for the x265_analysis_data object's internal structures. */
// void x265_alloc_analysis_data(x265_param *param, x265_analysis_data* analysis);
func (param *X265Param) X265AllocAnalysisData(analysis *X265AnalysisData) {
	libx265common.GetLibx265Dll().NewProc("x265_alloc_analysis_data").Call(
		uintptr(unsafe.Pointer(param)),
		uintptr(unsafe.Pointer(analysis)),
	)
}

// /*
// *    Free the allocated memory for x265_analysis_data object's internal structures. */
// void x265_free_analysis_data(x265_param *param, x265_analysis_data* analysis);

// /* Force a link error in the case of linking against an incompatible API version.
//  * Glue consts exist to force correct macro expansion; the final output of the macro
//  * is x265_encoder_open_##X265_BUILD (for purposes of dlopen). */
// const x265_encoder_glue1(x, y) x ## y
// const x265_encoder_glue2(x, y) x265_encoder_glue1(x, y)
// const x265_encoder_open x265_encoder_glue2(x265_encoder_open_, X265_BUILD)

// /* x265_encoder_open:
//   - create a new encoder handler, all parameters from x265_param are copied */
//
// x265_encoder* x265_encoder_open(x265_param *);
func (this *X265Param) X265EncoderOpen207() (res *X265Encoder) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_encoder_open_207").Call(
		uintptr(unsafe.Pointer(this)),
	)
	res = (*X265Encoder)(unsafe.Pointer(t))
	return
}

// /* x265_encoder_parameters:
//   - copies the current internal set of parameters to the pointer provided
//   - by the caller.  useful when the calling application needs to know
//   - how x265_encoder_open has changed the parameters.
//   - note that the data accessible through pointers in the returned param struct
//   - (e.g. filenames) should not be modified by the calling application. */
//
// void x265_encoder_parameters(x265_encoder *, x265_param *);
func (this *X265Encoder) X265EncoderParameters(p1 *X265Param) {
	libx265common.GetLibx265Dll().NewProc("x265_encoder_parameters").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
	)
}

// /* x265_encoder_headers:
//   - return the SPS and PPS that will be used for the whole stream.
//   - *pi_nal is the number of NAL units outputted in pp_nal.
//   - returns negative on error, total byte size of payload data on success
//   - the payloads of all output NALs are guaranteed to be sequential in memory. */
//
// int x265_encoder_headers(x265_encoder *, x265_nal **pp_nal, uint32_t *pi_nal);
func (this *X265Encoder) X265EncoderHeaders(pp_nal **X265Nal, pi_nal *ffcommon.FUint32T) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_encoder_headers").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pp_nal)),
		uintptr(unsafe.Pointer(pi_nal)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_encoder_encode:
//   - encode one picture.
//   - *pi_nal is the number of NAL units outputted in pp_nal.
//   - returns negative on error, 1 if a picture and access unit were output,
//   - or zero if the encoder pipeline is still filling or is empty after flushing.
//   - the payloads of all output NALs are guaranteed to be sequential in memory.
//   - To flush the encoder and retrieve delayed output pictures, pass pic_in as NULL.
//   - Once flushing has begun, all subsequent calls must pass pic_in as NULL. */
//
// int x265_encoder_encode(x265_encoder *encoder, x265_nal **pp_nal, uint32_t *pi_nal, x265_picture *pic_in, x265_picture *pic_out);
func (encoder *X265Encoder) X265EncoderEncode(pp_nal **X265Nal, pi_nal *ffcommon.FUint32T, pic_in, pic_out *X265Picture) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_encoder_encode").Call(
		uintptr(unsafe.Pointer(encoder)),
		uintptr(unsafe.Pointer(pp_nal)),
		uintptr(unsafe.Pointer(pi_nal)),
		uintptr(unsafe.Pointer(pic_in)),
		uintptr(unsafe.Pointer(pic_out)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_encoder_reconfig:
//   - various parameters from x265_param are copied.
//   - this takes effect immediately, on whichever frame is encoded next;
//   - returns 0 on success, negative on parameter validation error.
//     *
//   - not all parameters can be changed; see the actual function for a
//   - detailed breakdown.  since not all parameters can be changed, moving
//   - from preset to preset may not always fully copy all relevant parameters,
//   - but should still work usably in practice. however, more so than for
//   - other presets, many of the speed shortcuts used in ultrafast cannot be
//   - switched out of; using reconfig to switch between ultrafast and other
//   - presets is not recommended without a more fine-grained breakdown of
//   - parameters to take this into account. */
//
// int x265_encoder_reconfig(x265_encoder *, x265_param *);
func (this *X265Encoder) X265EncoderReconfig(p1 *X265Param) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_encoder_reconfig").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_encoder_reconfig_zone:
// *       zone settings are copied to the encoder's param.
// *       Properties of the zone will be used only to re-configure rate-control settings
// *       of the zone mid-encode. Returns 0 on success on successful copy, negative on failure.*/
// int x265_encoder_reconfig_zone(x265_encoder *, x265_zone *);
func (this *X265Encoder) X265EncoderReconfigZone(p1 *X265Zone) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_encoder_reconfig_zone").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_encoder_get_stats:
//   - returns encoder statistics */
//
// void x265_encoder_get_stats(x265_encoder *encoder, x265_stats *, uint32_t statsSizeBytes);
func (this *X265Encoder) X265EncoderGetStats(p1 *X265Stats, statsSizeBytes ffcommon.FUint32T) {
	libx265common.GetLibx265Dll().NewProc("x265_encoder_get_stats").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
		uintptr(statsSizeBytes),
	)
}

// /* x265_encoder_log:
//   - write a line to the configured CSV file.  If a CSV filename was not
//   - configured, or file open failed, this function will perform no write. */
//
// void x265_encoder_log(x265_encoder *encoder, int argc, char **argv);
func (encoder *X265Encoder) X265EncoderLog(args []string) {
	argc := len(args)
	argv := uintptr(0)
	if argc > 0 {
		a := make([]uintptr, argc, argc)
		for i := 0; i < argc; i++ {
			a[i] = ffcommon.UintPtrFromString(args[i])
		}
		argv = uintptr(unsafe.Pointer(&a[0]))
	}

	libx265common.GetLibx265Dll().NewProc("x265_encoder_log").Call(
		uintptr(unsafe.Pointer(encoder)),
		uintptr(argc),
		argv,
	)
}

// /* x265_encoder_close:
//   - close an encoder handler */
//
// void x265_encoder_close(x265_encoder *);
func (this *X265Encoder) X265EncoderClose() {
	libx265common.GetLibx265Dll().NewProc("x265_encoder_close").Call(
		uintptr(unsafe.Pointer(this)),
	)
}

// /* x265_encoder_intra_refresh:
//  *      If an intra refresh is not in progress, begin one with the next P-frame.
//  *      If an intra refresh is in progress, begin one as soon as the current one finishes.
//  *      Requires bIntraRefresh to be set.
//  *
//  *      Useful for interactive streaming where the client can tell the server that packet loss has
//  *      occurred.  In this case, keyint can be set to an extremely high value so that intra refreshes
//  *      occur only when calling x265_encoder_intra_refresh.
//  *
//  *      In multi-pass encoding, if x265_encoder_intra_refresh is called differently in each pass,
//  *      behavior is undefined.
//  *
//  *      Should not be called during an x265_encoder_encode. */

// int x265_encoder_intra_refresh(x265_encoder *);
func (this *X265Encoder) X265EncoderIntraRefresh() (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_encoder_intra_refresh").Call(
		uintptr(unsafe.Pointer(this)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_encoder_ctu_info:
//   - Copy CTU information such as ctu address and ctu partition structure of all
//   - CTUs in each frame. The function is invoked only if "--ctu-info" is enabled and
//   - the encoder will wait for this copy to complete if enabled.
//     */
//
// int x265_encoder_ctu_info(x265_encoder *, int poc, x265_ctu_info_t** ctu);
func (this *X265Encoder) X265EncoderCtuInfo(poc ffcommon.FInt, ctu **X265CtuInfoT) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_encoder_ctu_info").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(poc),
		uintptr(unsafe.Pointer(ctu)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_get_slicetype_poc_and_scenecut:
//   - get the slice type, poc and scene cut information for the current frame,
//   - returns negative on error, 0 when access unit were output.
//   - This API must be called after(poc >= lookaheadDepth + bframes + 2) condition check */
//
// int x265_get_slicetype_poc_and_scenecut(x265_encoder *encoder, int *slicetype, int *poc, int* sceneCut);
func (encoder *X265Encoder) X265GetSlicetypePocAndScenecut(slicetype, poc, sceneCut *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_get_slicetype_poc_and_scenecut").Call(
		uintptr(unsafe.Pointer(encoder)),
		uintptr(unsafe.Pointer(slicetype)),
		uintptr(unsafe.Pointer(poc)),
		uintptr(unsafe.Pointer(sceneCut)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_get_ref_frame_list:
//   - returns negative on error, 0 when access unit were output.
//   - This API must be called after(poc >= lookaheadDepth + bframes + 2) condition check */
//
// int x265_get_ref_frame_list(x265_encoder *encoder, x265_picyuv**, x265_picyuv**, int, int, int*, int*);
func (encoder *X265Encoder) X265GetRefFrameList(p1, p2 **X265Picyuv, p3, p4 ffcommon.FInt, p5, p6 *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_get_ref_frame_list").Call(
		uintptr(unsafe.Pointer(encoder)),
		uintptr(unsafe.Pointer(p1)),
		uintptr(unsafe.Pointer(p2)),
		uintptr(p3),
		uintptr(p4),
		uintptr(unsafe.Pointer(p5)),
		uintptr(unsafe.Pointer(p6)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_set_analysis_data:
//   - set the analysis data. The incoming analysis_data structure is assumed to be AVC-sized blocks.
//   - returns negative on error, 0 access unit were output. */
//
// int x265_set_analysis_data(x265_encoder *encoder, x265_analysis_data *analysis_data, int poc, uint32_t cuBytes);
func (encoder *X265Encoder) X265SetAnalysisData(analysis_data *X265AnalysisData, poc ffcommon.FInt, cuBytes ffcommon.FUint32T) (res ffcommon.FInt) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_set_analysis_data").Call(
		uintptr(unsafe.Pointer(encoder)),
		uintptr(poc),
		uintptr(cuBytes),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x265_cleanup:
//   - release library static allocations, reset configured CTU size */
//
// void x265_cleanup(void);
func X265Cleanup() {
	libx265common.GetLibx265Dll().NewProc("x265_cleanup").Call()
}

// /* Open a CSV log file. On success it returns a file handle which must be passed
//   - to x265_csvlog_frame() and/or x265_csvlog_encode(). The file handle must be
//   - closed by the caller using fclose(). If csv-loglevel is 0, then no frame logging
//   - header is written to the file. This function will return NULL if it is unable
//   - to open the file for write or if it detects a structure size skew */
//
// FILE* x265_csvlog_open(const x265_param *);
func (this *X265Param) X265CsvlogOpen() (res *FILE) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_csvlog_open").Call(
		uintptr(unsafe.Pointer(this)),
	)
	res = (*FILE)(unsafe.Pointer(t))
	return
}

// /* Log frame statistics to the CSV file handle. csv-loglevel should have been non-zero
//   - in the call to x265_csvlog_open() if this function is called. */
//
// void x265_csvlog_frame(const x265_param *, const x265_picture *);
func (this *X265Param) X265CsvlogFrame(p1 *X265Picture) {
	libx265common.GetLibx265Dll().NewProc("x265_csvlog_frame").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
	)
}

// /* Log final encode statistics to the CSV file handle. 'argc' and 'argv' are
//   - intended to be command line arguments passed to the encoder. padx and pady are
//   - padding offsets for conformance and can be given from sps settings. Encode
//   - statistics should be queried from the encoder just prior to closing it. */
//
// void x265_csvlog_encode(const x265_param*, const x265_stats *, int padx, int pady, int argc, char** argv);
func (this *X265Param) X265CsvlogEncode(p1 *X265Stats, padx, pady ffcommon.FInt, args []string) {
	argc := len(args)
	argv := uintptr(0)
	if argc > 0 {
		a := make([]uintptr, argc, argc)
		for i := 0; i < argc; i++ {
			a[i] = ffcommon.UintPtrFromString(args[i])
		}
		argv = uintptr(unsafe.Pointer(&a[0]))
	}
	libx265common.GetLibx265Dll().NewProc("x265_csvlog_encode").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
		uintptr(padx),
		uintptr(pady),
		uintptr(argc),
		argv,
	)
}

// /* In-place downshift from a bit-depth greater than 8 to a bit-depth of 8, using
//   - the residual bits to dither each row. */
//
// void x265_dither_image(x265_picture *, int picWidth, int picHeight, int16_t *errorBuf, int bitDepth);
func (this *X265Picture) X265DitherImage(picWidth, picHeight ffcommon.FInt, errorBuf *ffcommon.FInt16T, bitDepth ffcommon.FInt) {
	libx265common.GetLibx265Dll().NewProc("x265_dither_image").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(picWidth),
		uintptr(picHeight),
		uintptr(unsafe.Pointer(errorBuf)),
		uintptr(bitDepth),
	)
}

// #if ENABLE_LIBVMAF
// /* x265_calculate_vmafScore:
//   - returns VMAF score for the input video.
//   - This api must be called only after encoding was done. */
//
// double x265_calculate_vmafscore(x265_param*, x265_vmaf_data*);
func (this *X265Param) x265_calculate_vmafscore(p1 *X265VmafData) (res ffcommon.FDouble) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_calculate_vmafscore").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
	)
	res = *(*ffcommon.FDouble)(unsafe.Pointer(t))
	return
}

// /* x265_calculate_vmaf_framelevelscore:
//   - returns VMAF score for each frame in a given input video. */
//
// double x265_calculate_vmaf_framelevelscore(x265_vmaf_framedata*);
func (this *X265Param) X265CalculateVmafFramelevelscore() (res ffcommon.FDouble) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_calculate_vmaf_framelevelscore").Call(
		uintptr(unsafe.Pointer(this)),
	)
	res = *(*ffcommon.FDouble)(unsafe.Pointer(t))
	return
}

// /* x265_vmaf_encoder_log:
//   - write a line to the configured CSV file.  If a CSV filename was not
//   - configured, or file open failed, this function will perform no write.
//   - This api will be called only when ENABLE_LIBVMAF cmake option is set */
//
// void x265_vmaf_encoder_log(x265_encoder *encoder, int argc, char **argv, x265_param*, x265_vmaf_data*);
func (encoder *X265Encoder) X265VmafEncoderLog(args []string, p1 *X265Param, p2 *X265VmafData) {
	argc := len(args)
	argv := uintptr(0)
	if argc > 0 {
		a := make([]uintptr, argc, argc)
		for i := 0; i < argc; i++ {
			a[i] = ffcommon.UintPtrFromString(args[i])
		}
		argv = uintptr(unsafe.Pointer(&a[0]))
	}
	libx265common.GetLibx265Dll().NewProc("x265_vmaf_encoder_log").Call(
		uintptr(unsafe.Pointer(encoder)),
		uintptr(argc),
		argv,
		uintptr(unsafe.Pointer(p1)),
		uintptr(unsafe.Pointer(p2)),
	)
}

// #endif

const X265_MAJOR_VERSION = 1

/* === Multi-lib API ===
 * By using this method to gain access to the libx265 interfaces, you allow run-
 * time selection between various available libx265 libraries based on the
 * encoder parameters. The most likely use case is to choose between Main and
 * Main10 builds of libx265. */

type X265Api struct {
	ApiMajorVersion    ffcommon.FInt /* X265_MAJOR_VERSION */
	ApiBuildNumber     ffcommon.FInt /* X265_BUILD (soname) */
	SizeofParam        ffcommon.FInt /* sizeof(x265_param) */
	SizeofPicture      ffcommon.FInt /* sizeof(x265_picture) */
	SizeofAnalysisData ffcommon.FInt /* sizeof(x265_analysis_data) */
	SizeofZone         ffcommon.FInt /* sizeof(x265_zone) */
	SizeofStats        ffcommon.FInt /* sizeof(x265_stats) */

	Bit_depth      ffcommon.FInt
	Version_str    ffcommon.FCharPStruct
	Build_info_str ffcommon.FCharPStruct

	/* libx265 public API functions, documented above with x265_ prefixes */
	//x265_param*   (*param_alloc)(void);
	Param_alloc uintptr
	//void          (*param_free)(x265_param*);
	Param_free uintptr
	//void          (*param_default)(x265_param*);
	Param_default uintptr
	//int           (*param_parse)(x265_param*, const char*, const char*);
	Param_parse uintptr
	//int           (*scenecut_aware_qp_param_parse)(x265_param*, const char*, const char*);
	Scenecut_aware_qp_param_parse uintptr
	//int           (*param_apply_profile)(x265_param*, const char*);
	Param_apply_profile uintptr
	//int           (*param_default_preset)(x265_param*, const char*, const char *);
	Param_default_preset uintptr
	//x265_picture* (*picture_alloc)(void);
	Picture_alloc uintptr
	//void          (*picture_free)(x265_picture*);
	Picture_free uintptr
	//void          (*picture_init)(x265_param*, x265_picture*);
	Picture_init uintptr
	//x265_encoder* (*encoder_open)(x265_param*);
	Encoder_open uintptr
	//void          (*encoder_parameters)(x265_encoder*, x265_param*);
	Encoder_parameters uintptr
	//int           (*encoder_reconfig)(x265_encoder*, x265_param*);
	Encoder_reconfig uintptr
	//int           (*encoder_reconfig_zone)(x265_encoder*, x265_zone*);
	Encoder_reconfig_zone uintptr
	//int           (*encoder_headers)(x265_encoder*, x265_nal**, uint32_t*);
	Encoder_headers uintptr
	//int           (*encoder_encode)(x265_encoder*, x265_nal**, uint32_t*, x265_picture*, x265_picture*);
	Encoder_encode uintptr
	//void          (*encoder_get_stats)(x265_encoder*, x265_stats*, uint32_t);
	Encoder_get_stats uintptr
	//void          (*encoder_log)(x265_encoder*, int, char**);
	Encoder_log uintptr
	//void          (*encoder_close)(x265_encoder*);
	Encoder_close uintptr
	//void          (*cleanup)(void);
	Cleanup uintptr

	Sizeof_frame_stats ffcommon.FInt /* sizeof(x265_frame_stats) */
	//int           (*encoder_intra_refresh)(x265_encoder*);
	Encoder_intra_refresh uintptr
	//int           (*encoder_ctu_info)(x265_encoder*, int, x265_ctu_info_t**);
	Encoder_ctu_info uintptr
	//int           (*get_slicetype_poc_and_scenecut)(x265_encoder*, int*, int*, int*);
	Get_slicetype_poc_and_scenecut uintptr
	//int           (*get_ref_frame_list)(x265_encoder*, x265_picyuv**, x265_picyuv**, int, int, int*, int*);
	Get_ref_frame_list uintptr
	//FILE*         (*csvlog_open)(const x265_param*);
	Csvlog_open uintptr
	//void          (*csvlog_frame)(const x265_param*, const x265_picture*);
	Csvlog_frame uintptr
	//void          (*csvlog_encode)(const x265_param*, const x265_stats *, int, int, int, char**);
	Csvlog_encode uintptr
	//void          (*dither_image)(x265_picture*, int, int, int16_t*, int);
	Dither_image uintptr
	//int           (*set_analysis_data)(x265_encoder *encoder, x265_analysis_data *analysis_data, int poc, uint32_t cuBytes);
	Set_analysis_data uintptr
	// #if ENABLE_LIBVMAF
	// // double        (*calculate_vmafscore)(x265_param *, x265_vmaf_data *);
	// calculate_vmafscore uintptr
	// // double        (*calculate_vmaf_framelevelscore)(x265_vmaf_framedata *);
	// calculate_vmaf_framelevelscore uintptr
	// // void          (*vmaf_encoder_log)(x265_encoder*, int, char**, x265_param *, x265_vmaf_data *);
	// vmaf_encoder_log uintptr
	// #endif
	//int           (*zone_param_parse)(x265_param*, const char*, const char*);
	Zone_param_parse uintptr
	/* add new pointers to the end, or increment X265_MAJOR_VERSION */
}

// /* Force a link error in the case of linking against an incompatible API version.
//  * Glue consts exist to force correct macro expansion; the final output of the macro
//  * is x265_api_get_##X265_BUILD (for purposes of dlopen). */
// const x265_api_glue1(x, y) x ## y
// const x265_api_glue2(x, y) x265_api_glue1(x, y)
// const x265_api_get x265_api_glue2(x265_api_get_, X265_BUILD)

// /* x265_api_get:
//   - Retrieve the programming interface for a linked x265 library.
//   - May return NULL if no library is available that supports the
//   - requested bit depth. If bitDepth is 0 the function is guarunteed
//   - to return a non-NULL x265_api pointer, from the linked libx265.
//     *
//   - If the requested bitDepth is not supported by the linked libx265,
//   - it will attempt to dynamically bind x265_api_get() from a shared
//   - library with an appropriate name:
//   - 8bit:  libx265_main.so
//   - 10bit: libx265_main10.so
//   - Obviously the shared library file extension is platform specific */
//
// const x265_api* x265_api_get(int bitDepth);
func X265ApiGet207(bitDepth ffcommon.FInt) (res *X265Api) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_api_get_207").Call(
		uintptr(bitDepth),
	)
	res = (*X265Api)(unsafe.Pointer(t))
	return
}

// /* x265_api_query:
//   - Retrieve the programming interface for a linked x265 library, like
//   - x265_api_get(), except this function accepts X265_BUILD as the second
//   - argument rather than using the build number as part of the function name.
//   - Applications which dynamically link to libx265 can use this interface to
//   - query the library API and achieve a relative amount of version skew
//   - flexibility. The function may return NULL if the library determines that
//   - the apiVersion that your application was compiled against is not compatible
//   - with the library you have linked with.
//     *
//   - api_major_version will be incremented any time non-backward compatible
//   - changes are made to any public structures or functions. If
//   - api_major_version does not match X265_MAJOR_VERSION from the x265.h your
//   - application compiled against, your application must not use the returned
//   - x265_api pointer.
//     *
//   - Users of this API *must* also validate the sizes of any structures which
//   - are not treated as opaque in application code. For instance, if your
//   - application dereferences a x265_param pointer, then it must check that
//   - api->sizeof_param matches the sizeof(x265_param) that your application
//   - compiled with. */
//
// const x265_api* x265_api_query(int bitDepth, int apiVersion, int* err);
func X265ApiQuery(bitDepth, apiVersion ffcommon.FInt, err *ffcommon.FInt) (res *X265Api) {
	t, _, _ := libx265common.GetLibx265Dll().NewProc("x265_api_query").Call(
		uintptr(bitDepth),
		uintptr(apiVersion),
		uintptr(unsafe.Pointer(err)),
	)
	res = (*X265Api)(unsafe.Pointer(t))
	return
}

const X265_API_QUERY_ERR_NONE = 0           /* returned API pointer is non-NULL */
const X265_API_QUERY_ERR_VER_REFUSED = 1    /* incompatible version skew        */
const X265_API_QUERY_ERR_LIB_NOT_FOUND = 2  /* libx265_main10 not found, for ex */
const X265_API_QUERY_ERR_FUNC_NOT_FOUND = 3 /* unable to bind x265_api_query    */
const X265_API_QUERY_ERR_WRONG_BITDEPTH = 4 /* libx265_main10 not 10bit, for ex */

//	static const char * const x265_api_query_errnames[] = {
//	    "api queried from libx265",
//	    "libx265 version is not compatible with this application",
//	    "unable to bind a libx265 with requested bit depth",
//	    "unable to bind x265_api_query from libx265",
//	    "libx265 has an invalid bitdepth"
//	};
var X265ApiQueryErrnames = []string{"api queried from libx265",
	"libx265 version is not compatible with this application",
	"unable to bind a libx265 with requested bit depth",
	"unable to bind x265_api_query from libx265",
	"libx265 has an invalid bitdepth"}

// #ifdef __cplusplus
// }
// #endif

// #endif // X265_H
