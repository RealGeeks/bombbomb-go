package bombbomb_test

// Stub responses from the BombBomb API used in our tests

var Stubs = map[string]string{

	"AddContact": `{
		"status": "success",
		"methodName": "AddContact",
		"info": {
			"id": "106e0e29-e9cf-b812-b895-dcdc059cf9ec",
			"first_name": "Jack",
			"last_name": "Johnson",
			"email": "jj@gmail.com",
			"address_line_1": null,
			"address_line_2": null,
			"city": null,
			"state": null,
			"postal_code": null,
			"country": null,
			"phone_number": "808-123-4321",
			"business_name": null,
			"position": null,
			"opt_out": "0",
			"global_opt_out": "0",
			"last_bounce_date": null,
			"last_bounce_code": null,
			"bounce_details": null,
			"relationship_score": "0",
			"relationship_level_id": "1",
			"consent_status": "0",
			"relationship_level_name": "Weak",
			"last_engagement_date": null,
			"varchar_1": null,
			"varchar_2": null,
			"varchar_3": null,
			"varchar_4": null,
			"varchar_5": null,
			"varchar_6": null,
			"varchar_7": null,
			"varchar_8": null,
			"varchar_9": null,
			"varchar_10": null,
			"varchar_11": null,
			"varchar_12": null,
			"varchar_13": null,
			"varchar_14": null,
			"varchar_15": null,
			"varchar_16": null,
			"varchar_17": null,
			"varchar_18": null,
			"varchar_19": null,
			"varchar_20": null,
			"varchar_21": null,
			"varchar_22": null,
			"varchar_23": null,
			"varchar_24": null,
			"varchar_25": null,
			"varchar_26": null,
			"varchar_27": null,
			"varchar_28": null,
			"varchar_29": null,
			"varchar_30": null,
			"varchar_31": null,
			"varchar_32": null,
			"varchar_33": null,
			"varchar_34": null,
			"varchar_35": null,
			"varchar_36": null,
			"varchar_37": null,
			"varchar_38": null,
			"varchar_39": null,
			"varchar_40": null,
			"varchar_41": null,
			"varchar_42": null,
			"varchar_43": null,
			"varchar_44": null,
			"varchar_45": null,
			"varchar_46": null,
			"varchar_47": null,
			"varchar_48": null,
			"varchar_49": null,
			"varchar_50": null,
			"varchar_51": null,
			"varchar_52": null,
			"varchar_53": null,
			"varchar_54": null,
			"varchar_55": null,
			"varchar_56": null,
			"varchar_57": null,
			"varchar_58": null,
			"varchar_59": null,
			"varchar_60": null,
			"varchar_61": null,
			"varchar_62": null,
			"varchar_63": null,
			"varchar_64": null,
			"varchar_65": null,
			"varchar_66": null,
			"varchar_67": null,
			"varchar_68": null,
			"varchar_69": null,
			"varchar_70": null,
			"varchar_71": null,
			"varchar_72": null,
			"varchar_73": null,
			"varchar_74": null,
			"varchar_75": null,
			"varchar_76": null,
			"varchar_77": null,
			"varchar_78": null,
			"varchar_79": null,
			"varchar_80": null,
			"varchar_81": null,
			"varchar_82": null,
			"varchar_83": null,
			"varchar_84": null,
			"varchar_85": null,
			"varchar_86": null,
			"varchar_87": null,
			"varchar_88": null,
			"varchar_89": null,
			"varchar_90": null,
			"varchar_91": null,
			"varchar_92": null,
			"varchar_93": null,
			"varchar_94": null,
			"varchar_95": null,
			"varchar_96": null,
			"varchar_97": null,
			"varchar_98": null,
			"varchar_99": null,
			"varchar_100": null
		}
	}`,

	"CreateList": `{
		"status": "success",
		"methodName": "CreateList",
		"info": {
			"id": "4184993a-b98e-e9e4-19b6-da1019d9cd3d",
			"user_id": "668345da-b2c0-fe51-ce2b-a318dbce5865",
			"name": "Buyers",
			"description": null,
			"upload_date": null,
			"public_name": null,
			"agree_no_purchase_ts": null,
			"integrator_id": null,
			"integration_id": null,
			"integration_type": null,
			"pin": "0",
			"status": "0",
			"abuse_count": "0",
			"invalid_count": "0",
			"opt_out_count": "0",
			"sendable_count": "0",
			"send_count": "0",
			"import_status_message": null,
			"is_suppression": "0",
			"suppressed_count": "0",
			"oversubscribed_count": "0",
			"batch_send_list": "0",
			"automated_drip_id": null,
			"avg_relationship_score": "0",
			"relationship_level_id": "1",
			"status_updated": null,
			"avg_engagement": null,
			"last_send_date": null
		}
	}`,

	"GetLists": `{
		"status": "success",
		"methodName": "GetLists",
		"info": [{
			"id": "4184993a-b98e-e9e4-19b6-da1019d9cd3d",
			"name": "Partners",
			"ContactCount": "2"
		}, {
			"id": "3c20f8a3-2d95-8966-4add-0957dd0d23c5",
			"name": "Suppression List",
			"ContactCount": "0"
		}]
	}`,
}
