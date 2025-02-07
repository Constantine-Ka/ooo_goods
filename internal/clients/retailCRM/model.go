package retailCRM

type Order struct {
	OrderID            int
	CustomerID         int
	OrderNumber        string
	FactData           string
	PodtverjdenOtravka string
	TotalSumm          float64
	PrepaySumm         float64
}
type Response struct {
	Success    bool `json:"success"`
	Pagination struct {
		Limit          int `json:"limit"`
		TotalCount     int `json:"totalCount"`
		CurrentPage    int `json:"currentPage"`
		TotalPageCount int `json:"totalPageCount"`
	} `json:"pagination"`
	Orders []OrderResponse `json:"orders"`
}

type OrderResponse struct {
	Slug               int     `json:"slug"`
	BonusesCreditTotal int     `json:"bonusesCreditTotal"`
	BonusesChargeTotal int     `json:"bonusesChargeTotal"`
	Id                 int     `json:"id"`     //TODO
	Number             string  `json:"number"` //TODO
	OrderType          string  `json:"orderType"`
	OrderMethod        string  `json:"orderMethod"`
	PrivilegeType      string  `json:"privilegeType"`
	CreatedAt          string  `json:"createdAt"`
	StatusUpdatedAt    string  `json:"statusUpdatedAt"`
	Summ               int     `json:"summ"`
	TotalSumm          float64 `json:"totalSumm"` //TODO
	PrepaySum          float64 `json:"prepaySum"` //TODO
	PurchaseSumm       int     `json:"purchaseSumm"`
	MarkDatetime       string  `json:"markDatetime"`
	LastName           string  `json:"lastName"`
	FirstName          string  `json:"firstName"`
	Call               bool    `json:"call"`
	Expired            bool    `json:"expired"`
	ManagerId          int     `json:"managerId"`
	Customer           struct {
		Type       string `json:"type"`
		Id         int    `json:"id"` //TODO
		ExternalId string `json:"externalId"`
		IsContact  bool   `json:"isContact"`
		CreatedAt  string `json:"createdAt"`
		ManagerId  int    `json:"managerId"`
		Vip        bool   `json:"vip"`
		Bad        bool   `json:"bad"`
		Site       string `json:"site"`
		Contragent struct {
			ContragentType string `json:"contragentType"`
		} `json:"contragent"`
		Tags []struct {
			Color     string `json:"color"`
			Name      string `json:"name"`
			ColorCode string `json:"colorCode"`
			Attached  bool   `json:"attached"`
		} `json:"tags"`
		CustomFields struct {
			Hash                   string `json:"hash"`
			VkUrl                  string `json:"vk_url"`
			AskaChat               bool   `json:"aska_chat"`
			IdKlienta              int    `json:"id_klienta"`
			TestRassylki           bool   `json:"test_rassylki"`
			IdRassylkiKlient       int    `json:"id_rassylki_klient,omitempty"`
			SpisokIdRassylokKlient string `json:"spisok_id_rassylok_klient,omitempty"`
			Ref                    string `json:"ref,omitempty"`
			RefSource              string `json:"ref_source,omitempty"`
			RassylkaKlient03022025 bool   `json:"rassylka_klient03022025,omitempty"`
		} `json:"customFields"`
		PersonalDiscount int `json:"personalDiscount"`
		MarginSumm       int `json:"marginSumm"`
		TotalSumm        int `json:"totalSumm"`
		AverageSumm      int `json:"averageSumm"`
		OrdersCount      int `json:"ordersCount"`
		CostSumm         int `json:"costSumm"`
		Segments         []struct {
			Id             int    `json:"id"`
			Code           string `json:"code"`
			Name           string `json:"name"`
			CreatedAt      string `json:"createdAt"`
			IsDynamic      bool   `json:"isDynamic"`
			Active         bool   `json:"active"`
			CustomersCount int    `json:"customersCount,omitempty"`
		} `json:"segments"`
		FirstName             string        `json:"firstName"`
		LastName              string        `json:"lastName"`
		Sex                   string        `json:"sex,omitempty"`
		PresumableSex         string        `json:"presumableSex,omitempty"`
		CustomerSubscriptions []interface{} `json:"customerSubscriptions"`
		Phones                []interface{} `json:"phones"`
		MgCustomers           []struct {
			Id         int `json:"id"`
			ExternalId int `json:"externalId"`
			MgChannel  struct {
				AllowedSendByPhone bool   `json:"allowedSendByPhone"`
				Id                 int    `json:"id"`
				ExternalId         int    `json:"externalId"`
				Type               string `json:"type"`
				Active             bool   `json:"active"`
				Name               string `json:"name"`
			} `json:"mgChannel"`
		} `json:"mgCustomers"`
	} `json:"customer"`
	Contact struct {
		Type       string `json:"type"`
		Id         int    `json:"id"`
		ExternalId string `json:"externalId"`
		IsContact  bool   `json:"isContact"`
		CreatedAt  string `json:"createdAt"`
		ManagerId  int    `json:"managerId"`
		Vip        bool   `json:"vip"`
		Bad        bool   `json:"bad"`
		Site       string `json:"site"`
		Contragent struct {
			ContragentType string `json:"contragentType"`
		} `json:"contragent"`
		Tags []struct {
			Color     string `json:"color"`
			Name      string `json:"name"`
			ColorCode string `json:"colorCode"`
			Attached  bool   `json:"attached"`
		} `json:"tags"`
		CustomFields struct {
			Hash                   string `json:"hash"`
			VkUrl                  string `json:"vk_url"`
			AskaChat               bool   `json:"aska_chat"`
			IdKlienta              int    `json:"id_klienta"`
			TestRassylki           bool   `json:"test_rassylki"`
			IdRassylkiKlient       int    `json:"id_rassylki_klient,omitempty"`
			SpisokIdRassylokKlient string `json:"spisok_id_rassylok_klient,omitempty"`
			Ref                    string `json:"ref,omitempty"`
			RefSource              string `json:"ref_source,omitempty"`
			RassylkaKlient03022025 bool   `json:"rassylka_klient03022025,omitempty"`
		} `json:"customFields"`
		PersonalDiscount int `json:"personalDiscount"`
		MarginSumm       int `json:"marginSumm"`
		TotalSumm        int `json:"totalSumm"`
		AverageSumm      int `json:"averageSumm"`
		OrdersCount      int `json:"ordersCount"`
		CostSumm         int `json:"costSumm"`
		Segments         []struct {
			Id             int    `json:"id"`
			Code           string `json:"code"`
			Name           string `json:"name"`
			CreatedAt      string `json:"createdAt"`
			IsDynamic      bool   `json:"isDynamic"`
			Active         bool   `json:"active"`
			CustomersCount int    `json:"customersCount,omitempty"`
		} `json:"segments"`
		FirstName             string `json:"firstName"`
		LastName              string `json:"lastName"`
		Sex                   string `json:"sex,omitempty"`
		PresumableSex         string `json:"presumableSex,omitempty"`
		CustomerSubscriptions []struct {
			Subscription struct {
				Id            int    `json:"id"`
				Channel       string `json:"channel"`
				Name          string `json:"name"`
				Code          string `json:"code"`
				Active        bool   `json:"active"`
				AutoSubscribe bool   `json:"autoSubscribe"`
				Ordering      int    `json:"ordering"`
			} `json:"subscription"`
			Subscribed bool `json:"subscribed"`
		} `json:"customerSubscriptions"`
		Phones      []interface{} `json:"phones"`
		MgCustomers []struct {
			Id         int `json:"id"`
			ExternalId int `json:"externalId"`
			MgChannel  struct {
				AllowedSendByPhone bool   `json:"allowedSendByPhone"`
				Id                 int    `json:"id"`
				ExternalId         int    `json:"externalId"`
				Type               string `json:"type"`
				Active             bool   `json:"active"`
				Name               string `json:"name"`
			} `json:"mgChannel"`
		} `json:"mgCustomers"`
	} `json:"contact"`
	Contragent struct {
		ContragentType string `json:"contragentType"`
	} `json:"contragent"`
	Delivery struct {
		Code            string `json:"code"`
		IntegrationCode string `json:"integrationCode"`
		Data            struct {
			Locked          bool   `json:"locked"`
			PayerType       string `json:"payerType"`
			ShipmentpointId string `json:"shipmentpointId"`
			ExtraData       struct {
				ManualDeclareValue int           `json:"manual_declare_value"`
				Services           []interface{} `json:"services"`
			} `json:"extraData"`
			ItemDeclaredValues []interface{} `json:"itemDeclaredValues"`
			Packages           []interface{} `json:"packages"`
		} `json:"data"`
		Cost    int `json:"cost"`
		NetCost int `json:"netCost"`
		Address struct {
		} `json:"address"`
		VatRate string `json:"vatRate"`
	} `json:"delivery"`
	Site          string        `json:"site"`
	Status        string        `json:"status"`
	StatusComment string        `json:"statusComment"`
	Items         []interface{} `json:"items"`
	Payments      struct {
		Field1 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1719076,omitempty"`
		Field2 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1719048,omitempty"`
		Field3 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1719017,omitempty"`
		Field4 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1719000,omitempty"`
		Field5 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718999,omitempty"`
		Field6 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718969,omitempty"`
		Field7 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718948,omitempty"`
		Field8 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718926,omitempty"`
		Field9 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718870,omitempty"`
		Field10 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718828,omitempty"`
		Field11 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718760,omitempty"`
		Field12 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718684,omitempty"`
		Field13 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1718402,omitempty"`
		Field14 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1717756,omitempty"`
		Field15 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1717479,omitempty"`
		Field16 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1717313,omitempty"`
		Field17 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1717204,omitempty"`
		Field18 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1716751,omitempty"`
		Field19 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1716490,omitempty"`
		Field20 struct {
			Id     int    `json:"id"`
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"1715936,omitempty"`
	} `json:"payments"`
	FromApi      bool `json:"fromApi"`
	Shipped      bool `json:"shipped"`
	CustomFields struct {
		Box                         bool   `json:"box"`
		Fio                         string `json:"fio"`
		Rtu                         bool   `json:"rtu"`
		Chat                        string `json:"chat"`
		Assch                       string `json:"assch"`
		Blok1                       string `json:"blok1"`
		Blok2                       string `json:"blok2"`
		Blok4                       string `json:"blok4"`
		Blok5                       string `json:"blok5"`
		Otdel                       string `json:"otdel"`
		Blok41                      string `json:"blok41"`
		PoApi                       string `json:"po_api"`
		Rezerv                      bool   `json:"rezerv"`
		Askabot                     bool   `json:"askabot"`
		Doplata                     bool   `json:"doplata"`
		Poditog                     string `json:"poditog"`
		Prodaja                     string `json:"prodaja"`
		UrLico                      string `json:"ur_lico"`
		Vozvrat                     bool   `json:"vozvrat"`
		AskaConv                    bool   `json:"aska_conv"`
		FaktData                    string `json:"fakt_data"` //TODO
		Gravirovka                  bool   `json:"gravirovka"`
		KolVoFot                    int    `json:"kol_vo_fot"`
		MbarkIpsh                   bool   `json:"mbark_ipsh"`
		NozhCheap                   string `json:"nozh_cheap"`
		RndNumber                   int    `json:"rnd_number"`
		SmsSkidka                   bool   `json:"sms_skidka"`
		SummerSale                  bool   `json:"summer_sale"`
		Superadmins                 bool   `json:"superadmins"`
		UrlProfile                  string `json:"url_profile"`
		NetPodytoga                 bool   `json:"net_podytoga"`
		ChuzhoiZakaz                string `json:"chuzhoi_zakaz"`
		ObnovlenieBd                bool   `json:"obnovlenie_bd"`
		PostomatSdek                bool   `json:"postomat_sdek"`
		DefitsitSklad               bool   `json:"defitsit_sklad"`
		RekObnovlenie               bool   `json:"rek_obnovlenie"`
		SertifikatIsp               string `json:"sertifikat_isp"`
		SrochnyiZakaz               bool   `json:"srochnyi_zakaz"`
		OfferOtpravlen              string `json:"offer_otpravlen,omitempty"`
		SmenaRefMetki               bool   `json:"smena_ref_metki"`
		UpakovkaTovara              bool   `json:"upakovka_tovara"`
		IstochnikZakaza             string `json:"istochnik_zakaza"`
		Rassylka31072024            bool   `json:"rassylka31072024"`
		StarayaProdazha             string `json:"staraya_prodazha"`
		DefitsitVZakaze             bool   `json:"defitsit_v_zakaze"`
		PolKlientaZakaz             string `json:"pol_klienta_zakaz,omitempty"`
		UpsellTotalSumm             int    `json:"upsell_total_summ"`
		AktsiiSoSkidkami            string `json:"aktsii_so_skidkami"`
		NeDoshelDoTseny             bool   `json:"ne_doshel_do_tseny"`
		PosylkaPoteriana1           bool   `json:"posylka_poteriana1"`
		UmikoishiVZakaze            string `json:"umikoishi_v_zakaze"`
		KorobochkaVZakaze           string `json:"korobochka_v_zakaze"`
		ProverkaReklamacii          bool   `json:"proverka_reklamacii"`
		ReclamaciaCheckbox          bool   `json:"reclamacia_checkbox"`
		RezervDliaZakaza1           bool   `json:"rezerv_dlia_zakaza1"`
		VozvratPriniatV1S           bool   `json:"vozvrat_priniat_v1s"`
		DataSozdaniiaChata          string `json:"data_sozdaniia_chata"`
		Otrabotal2Mesiatsa1         string `json:"otrabotal2_mesiatsa1"`
		PereraspredelenieDs         string `json:"pereraspredelenie_ds"`
		ZakazBylVOtmenene           bool   `json:"zakaz_byl_v_otmenene"`
		DefitsitOt5TiNozhei         bool   `json:"defitsit_ot5ti_nozhei"`
		VremiaPostupleniiaZakaza    string `json:"vremia_postupleniia_zakaza"`
		NuzhenZvonokPoVrucheniiu    bool   `json:"nuzhen_zvonok_po_vrucheniiu"`
		NetDeneg                    string `json:"net_deneg,omitempty"`
		PotrebnostVyiavlena         string `json:"potrebnost_vyiavlena,omitempty"`
		TovarPrezentovanData        string `json:"tovar_prezentovan_data,omitempty"`
		KontaktUstanovlenData       string `json:"kontakt_ustanovlen_data,omitempty"`
		Ref                         string `json:"ref,omitempty"`
		RefSource                   string `json:"ref_source,omitempty"`
		Gipoteza4                   string `json:"gipoteza4,omitempty"`
		DorogoData                  string `json:"dorogo_data,omitempty"`
		PoiavilisVozrazheniiaData   string `json:"poiavilis_vozrazheniia_data,omitempty"`
		UtochnilVariantDostavki     string `json:"utochnil_variant_dostavki,omitempty"`
		PoluchilDannyeOtKlientaData string `json:"poluchil_dannye_ot_klienta_data,omitempty"`
		PodtverzdenOtpravka         string `json:"podtverzden_otpravka,omitempty"` //TODO
	} `json:"customFields"`
	Currency string `json:"currency"`
}
