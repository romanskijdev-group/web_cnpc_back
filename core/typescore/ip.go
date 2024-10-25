package typescore

type RegionInfoDetected struct {
	City        *string
	Region      *string
	CountryCode *string
	CountryName *string
}

type DetectorIPStruct struct {
	IP            *string
	IsINBlackList *bool
	RegionInfo    *RegionInfoDetected
}
