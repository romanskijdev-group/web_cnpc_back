package ipdetector

import (
	"cnpc_backend/core/typescore"
	"cnpc_backend/core/utilscore"
	"errors"
	"log"
	"strings"
)

// обработчик ip адреса
func (m *IPdetectorM) IpWorker(ip string) (*typescore.DetectorIPStruct, error) {
	objIpInfo := &typescore.DetectorIPStruct{}

	if strings.Contains(ip, ":") {
		parts := strings.Split(ip, ":")
		if len(parts) == 0 {
			err := errors.New("error split ip")
			log.Println("💔 error ipWorker ip", err)
			return nil, err
		}
		ip = parts[0]
		if len(ip) == 0 {
			err := errors.New("error split ip")
			log.Println("💔 error ipWorker ip", err)
			return nil, err
		}
	}

	ipM := &ip
	objIpInfo.IP = ipM
	err := m.checkerBlackListIps(ipM)
	if err != nil {
		objIpInfo.IsINBlackList = utilscore.PointerToBool(true)
	}
	objIpInfo.IsINBlackList = utilscore.PointerToBool(false)

	detectIp, err := m.IpPositionDetect(ip)
	if err != nil {
		log.Println("💔 error detect ip position", err)
	} else {
		objIpInfo.RegionInfo = detectIp
	}

	return objIpInfo, nil
}

// проверка ip адреса на наличие в черном списке
func (m *IPdetectorM) checkerBlackListIps(ip *string) error {
	errW := m.BlackListIP.CheckIPFromBlackList(ip)
	if errW != nil {
		log.Println("💔 error check ip from black list", errW)
		return errW.Err
	}
	return nil
}
