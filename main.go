package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

type Mediainfo struct {
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			TypeVideo               string `json:"@type"`
			Count                   string `json:"Count"`
			StreamCount             string `json:"StreamCount"`
			StreamKind              string `json:"StreamKind"`
			StreamKindString        string `json:"StreamKind_String"`
			StreamKindID            string `json:"StreamKindID"`
			VideoCount              string `json:"VideoCount,omitempty"`
			OtherCount              string `json:"OtherCount,omitempty"`
			VideoFormatList         string `json:"Video_Format_List,omitempty"`
			VideoFormatWithHintList string `json:"Video_Format_WithHint_List,omitempty"`
			VideoCodecList          string `json:"Video_Codec_List,omitempty"`
			CompleteName            string `json:"CompleteName,omitempty"`
			FileNameExtension       string `json:"FileNameExtension,omitempty"`
			FileName                string `json:"FileName,omitempty"`
			FileExtension           string `json:"FileExtension,omitempty"`
			Format                  string `json:"Format,omitempty"`
			FormatString            string `json:"Format_String,omitempty"`
			FormatExtensions        string `json:"Format_Extensions,omitempty"`
			FormatCommercial        string `json:"Format_Commercial,omitempty"`
			FormatProfile           string `json:"Format_Profile,omitempty"`
			InternetMediaType       string `json:"InternetMediaType,omitempty"`
			CodecID                 string `json:"CodecID,omitempty"`
			CodecIDString           string `json:"CodecID_String,omitempty"`
			CodecIDURL              string `json:"CodecID_Url,omitempty"`
			CodecIDVersion          string `json:"CodecID_Version,omitempty"`
			CodecIDCompatible       string `json:"CodecID_Compatible,omitempty"`
			FileSize                string `json:"FileSize,omitempty"`
			FileSizeString          string `json:"FileSize_String,omitempty"`
			FileSizeString1         string `json:"FileSize_String1,omitempty"`
			FileSizeString2         string `json:"FileSize_String2,omitempty"`
			FileSizeString3         string `json:"FileSize_String3,omitempty"`
			FileSizeString4         string `json:"FileSize_String4,omitempty"`
			Duration                string `json:"Duration"`
			DurationString          string `json:"Duration_String"`
			DurationString1         string `json:"Duration_String1"`
			DurationString2         string `json:"Duration_String2"`
			DurationString3         string `json:"Duration_String3"`
			DurationString4         string `json:"Duration_String4,omitempty"`
			DurationString5         string `json:"Duration_String5"`
			OverallBitRate          string `json:"OverallBitRate,omitempty"`
			OverallBitRateString    string `json:"OverallBitRate_String,omitempty"`
			FrameRate               string `json:"FrameRate,omitempty"`
			FrameRateString         string `json:"FrameRate_String,omitempty"`
			FrameCount              string `json:"FrameCount"`
			StreamSize              string `json:"StreamSize,omitempty"`
			StreamSizeString        string `json:"StreamSize_String,omitempty"`
			StreamSizeString1       string `json:"StreamSize_String1,omitempty"`
			StreamSizeString2       string `json:"StreamSize_String2,omitempty"`
			StreamSizeString3       string `json:"StreamSize_String3,omitempty"`
			StreamSizeString4       string `json:"StreamSize_String4,omitempty"`
			StreamSizeString5       string `json:"StreamSize_String5,omitempty"`
			StreamSizeProportion    string `json:"StreamSize_Proportion,omitempty"`
			HeaderSize              string `json:"HeaderSize,omitempty"`
			DataSize                string `json:"DataSize,omitempty"`
			FooterSize              string `json:"FooterSize,omitempty"`
			IsStreamable            string `json:"IsStreamable,omitempty"`
			EncodedDate             string `json:"Encoded_Date,omitempty"`
			TaggedDate              string `json:"Tagged_Date,omitempty"`
			FileModifiedDate        string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal   string `json:"File_Modified_Date_Local,omitempty"`
			EncodedLibrary          string `json:"Encoded_Library,omitempty"`
			EncodedLibraryString    string `json:"Encoded_Library_String,omitempty"`
			EncodedLibraryName      string `json:"Encoded_Library_Name,omitempty"`
			ExtraVideo              struct {
				ComAppleQuicktimeLocationAccuracyHorizontal string `json:"com_apple_quicktime_location_accuracy_horizontal"`
				ComAppleQuicktimeLocationISO6709            string `json:"com_apple_quicktime_location_ISO6709"`
				ComAppleQuicktimeMake                       string `json:"com_apple_quicktime_make"`
				ComAppleQuicktimeModel                      string `json:"com_apple_quicktime_model"`
				ComAppleQuicktimeSoftware                   string `json:"com_apple_quicktime_software"`
				ComAppleQuicktimeCreationdate               string `json:"com_apple_quicktime_creationdate"`
			} `json:"extra,omitempty"`
			StreamOrder                    string `json:"StreamOrder,omitempty"`
			ID                             string `json:"ID,omitempty"`
			IDString                       string `json:"ID_String,omitempty"`
			FormatInfo                     string `json:"Format_Info,omitempty"`
			FormatURL                      string `json:"Format_Url,omitempty"`
			FormatLevel                    string `json:"Format_Level,omitempty"`
			FormatSettings                 string `json:"Format_Settings,omitempty"`
			FormatSettingsCABAC            string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsCABACString      string `json:"Format_Settings_CABAC_String,omitempty"`
			FormatSettingsRefFrames        string `json:"Format_Settings_RefFrames,omitempty"`
			FormatSettingsRefFramesString  string `json:"Format_Settings_RefFrames_String,omitempty"`
			FormatSettingsGOP              string `json:"Format_Settings_GOP,omitempty"`
			CodecIDInfo                    string `json:"CodecID_Info,omitempty"`
			BitRate                        string `json:"BitRate,omitempty"`
			BitRateString                  string `json:"BitRate_String,omitempty"`
			Width                          string `json:"Width,omitempty"`
			WidthString                    string `json:"Width_String,omitempty"`
			Height                         string `json:"Height,omitempty"`
			HeightString                   string `json:"Height_String,omitempty"`
			StoredWidth                    string `json:"Stored_Width,omitempty"`
			SampledWidth                   string `json:"Sampled_Width,omitempty"`
			SampledHeight                  string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
			DisplayAspectRatioString       string `json:"DisplayAspectRatio_String,omitempty"`
			Rotation                       string `json:"Rotation,omitempty"`
			RotationString                 string `json:"Rotation_String,omitempty"`
			FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
			FrameRateModeString            string `json:"FrameRate_Mode_String,omitempty"`
			FrameRateNum                   string `json:"FrameRate_Num,omitempty"`
			FrameRateDen                   string `json:"FrameRate_Den,omitempty"`
			FrameRateMinimum               string `json:"FrameRate_Minimum,omitempty"`
			FrameRateMinimumString         string `json:"FrameRate_Minimum_String,omitempty"`
			FrameRateMaximum               string `json:"FrameRate_Maximum,omitempty"`
			FrameRateMaximumString         string `json:"FrameRate_Maximum_String,omitempty"`
			ColorSpace                     string `json:"ColorSpace,omitempty"`
			ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
			ChromaSubsamplingString        string `json:"ChromaSubsampling_String,omitempty"`
			BitDepth                       string `json:"BitDepth,omitempty"`
			BitDepthString                 string `json:"BitDepth_String,omitempty"`
			ScanType                       string `json:"ScanType,omitempty"`
			ScanTypeString                 string `json:"ScanType_String,omitempty"`
			BitsPixelFrame                 string `json:"BitsPixel_Frame,omitempty"`
			Title                          string `json:"Title,omitempty"`
			ColourDescriptionPresent       string `json:"colour_description_present,omitempty"`
			ColourDescriptionPresentSource string `json:"colour_description_present_Source,omitempty"`
			ColourRange                    string `json:"colour_range,omitempty"`
			ColourRangeSource              string `json:"colour_range_Source,omitempty"`
			ColourPrimaries                string `json:"colour_primaries,omitempty"`
			ColourPrimariesSource          string `json:"colour_primaries_Source,omitempty"`
			TransferCharacteristics        string `json:"transfer_characteristics,omitempty"`
			TransferCharacteristicsSource  string `json:"transfer_characteristics_Source,omitempty"`
			MatrixCoefficients             string `json:"matrix_coefficients,omitempty"`
			MatrixCoefficientsSource       string `json:"matrix_coefficients_Source,omitempty"`
			Extra                          struct {
				CodecConfigurationBox string `json:"CodecConfigurationBox"`
			} `json:"extra,omitempty"`
			Typeorder     string `json:"@typeorder,omitempty"`
			StreamKindPos string `json:"StreamKindPos,omitempty"`
			Type          string `json:"Type,omitempty"`
		} `json:"track"`
	} `json:"media"`
}


func main() {


	router := mux.NewRouter()
	router.HandleFunc("/", Redirect).Methods("GET")
	router.HandleFunc("/healthcheck", GetWorking).Methods("GET")
	router.HandleFunc("/videoinfo/{id}", VideoInfo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/healthcheck", http.StatusFound)
}

func GetWorking(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "WORKING")
}

func VideoInfo(w http.ResponseWriter, r *http.Request) {

	execMedia := exec.Command("mediainfo", "--Output=JSON", "nomedoarquivo")

	out, _ := execMedia.CombinedOutput()

	var e Mediainfo

	json.Unmarshal(out, &e)

	fmt.Print(e)

	json.NewEncoder(w).Encode(e)

}
