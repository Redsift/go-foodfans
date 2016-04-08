// Code generated by "stringer -type=Verb,Adjective,Noun"; DO NOT EDIT

package lookup

import "fmt"

const _Verb_name = "arrangebastebeatblendbrownburycarvecheckchopclosecovercrumplecutdecoratediscarddividedrapedropfilmfoldfollowformforceglazeinsertlayleaveliftmakemeltmincemixmoistenmoundopenpackpaintpiercepourpreparepressprickpullpureepushquarterraisereducerefreshreheatreplaceringroastrollsaltsautescatterscoopscrapescrubseasonseparatesetsettleshavesimmerskimsliceslideslipslitsmearsoakspoonspreadsprinklestirstrainstrewstuffsurroundtastetietilttiptoptosstrimturntwistwiltwindwrapLastVerb"

var _Verb_index = [...]uint16{0, 7, 12, 16, 21, 26, 30, 35, 40, 44, 49, 54, 61, 64, 72, 79, 85, 90, 94, 98, 102, 108, 112, 117, 122, 128, 131, 136, 140, 144, 148, 153, 156, 163, 168, 172, 176, 181, 187, 191, 198, 203, 208, 212, 217, 221, 228, 233, 239, 246, 252, 259, 263, 268, 272, 276, 281, 288, 293, 299, 304, 310, 318, 321, 327, 332, 338, 342, 347, 352, 356, 360, 365, 369, 374, 380, 388, 392, 398, 403, 408, 416, 421, 424, 428, 431, 434, 438, 442, 446, 451, 455, 459, 463, 471}

func (i Verb) String() string {
	if i < 0 || i >= Verb(len(_Verb_index)-1) {
		return fmt.Sprintf("Verb(%d)", i)
	}
	return _Verb_name[_Verb_index[i]:_Verb_index[i+1]]
}

const _Adjective_name = "acerbicacidicacridagedambrosialampleappealingappetizingaromaticastringentbakedbalsamicbeautifulbite_sizebitterblandblendedboiledbrackishbrinybrownedburntbutteredcakedcandiedcaramelizedcausticcenter_cutchar_broiledcheesychilledchoicecholesterol_freechunkedclassicclassyclove_coatedcoldcoolcopiouscountrycraftedcreamedcreamycrispcrunchycureddazzlingdeep_frieddelectabledeliciousdelightfuldistinctivedivinedoughydresseddrippingdrizzleddrydulcifieddulledibleelasticencrustedethnicextraordinaryfamousfantasticfetidfieryfiletfizzyflakyflatflavorfulflavorlessflavorsomefleshyfluffyfragilefreefree_rangefreshfriedfrostyfrozenfruityfullfull_bodiedfurrygarlickygenerousgingeryglazedgoldengorgeousgourmetgreasygrilledgrittygustatoryhalfharshheadyheapingheart_healthyheart_smartheartyheavenlyhomemadehoneyedhoney_glazedhotice_coldicyincisiveindulgentinfusedinsipidintenseintriguingjuicyjumbokosherlargelavishlayeredleanleatherylemonlightlitelightly_saltedlightly_breadedlip_smackinglivelylowlow_sodiumlow_fatlukewarmlusciouslushmarinatedmashedmellowmildmintymixedmoistmouth_wateringinternationally_famousnationally_famousnaturalnectarousnon_fatnuttyoilyorganicoverpoweringpepperypetitepickledpiquantplainpleasantplumppoachedpopularpoundedpreparedpulpypungentpureedrancidrankreducedrichriperoastedrobustrottenrubberysaccharinesalinesaltysapidsaporificsaporoussatinsatinysauteedsavorlesssavoryscrumptioussea_saltsearedseasonedsharp_tastingsilkysimmeredsizzlingskillfullysmallsmellysmokedsmokysmoothsmotheredsoothingsoursouthern_stylespecialspicedspicyspiral_cutspongysprinkledstalesteamedsteamystickystingingstrawberry_flavoredstrongstuffedsucculentsugar_coatedsugar_freesugaredsugarlesssugarysuperbsweetsweet_and_soursweetenedsyrupytangytantalizingtarttastefultastelesstastytepidterrificthickthintoastedtoothsometoppedtossedtoughtraditionaltreaclytreatunflavoredunsavoryunseasonedvanilla_flavoredvelvetyvinegarywarmwaxyweakwhippedwholewonderfulyuckyyummyzestyzingyboringdreamydrunkgoofynaughtypricklysharpsuspicioustenderLastAdjective"

var _Adjective_index = [...]uint16{0, 7, 13, 18, 22, 31, 36, 45, 55, 63, 73, 78, 86, 95, 104, 110, 115, 122, 128, 136, 141, 148, 153, 161, 166, 173, 184, 191, 201, 213, 219, 226, 232, 248, 255, 262, 268, 280, 284, 288, 295, 302, 309, 316, 322, 327, 334, 339, 347, 357, 367, 376, 386, 397, 403, 409, 416, 424, 432, 435, 444, 448, 454, 461, 470, 476, 489, 495, 504, 509, 514, 519, 524, 529, 533, 542, 552, 562, 568, 574, 581, 585, 595, 600, 605, 611, 617, 623, 627, 638, 643, 651, 659, 666, 672, 678, 686, 693, 699, 706, 712, 721, 725, 730, 735, 742, 755, 766, 772, 780, 788, 795, 807, 810, 818, 821, 829, 838, 845, 852, 859, 869, 874, 879, 885, 890, 896, 903, 907, 915, 920, 925, 929, 943, 958, 970, 976, 979, 989, 996, 1004, 1012, 1016, 1025, 1031, 1037, 1041, 1046, 1051, 1056, 1070, 1092, 1109, 1116, 1125, 1132, 1137, 1141, 1148, 1160, 1167, 1173, 1180, 1187, 1192, 1200, 1205, 1212, 1219, 1226, 1234, 1239, 1246, 1252, 1258, 1262, 1269, 1273, 1277, 1284, 1290, 1296, 1303, 1313, 1319, 1324, 1329, 1338, 1346, 1351, 1357, 1364, 1373, 1379, 1390, 1398, 1404, 1412, 1425, 1430, 1438, 1446, 1456, 1461, 1467, 1473, 1478, 1484, 1493, 1501, 1505, 1519, 1526, 1532, 1537, 1547, 1553, 1562, 1567, 1574, 1580, 1586, 1594, 1613, 1619, 1626, 1635, 1647, 1657, 1664, 1673, 1679, 1685, 1690, 1704, 1713, 1719, 1724, 1735, 1739, 1747, 1756, 1761, 1766, 1774, 1779, 1783, 1790, 1799, 1805, 1811, 1816, 1827, 1834, 1839, 1849, 1857, 1867, 1883, 1890, 1898, 1902, 1906, 1910, 1917, 1922, 1931, 1936, 1941, 1946, 1951, 1957, 1963, 1968, 1973, 1980, 1987, 1992, 2002, 2008, 2021}

func (i Adjective) String() string {
	if i < 0 || i >= Adjective(len(_Adjective_index)-1) {
		return fmt.Sprintf("Adjective(%d)", i)
	}
	return _Adjective_name[_Adjective_index[i]:_Adjective_index[i+1]]
}

const _Noun_name = "asparagusappleapricotanchoviesapplesaucealphabits_cerealartichokesanimal_crackersalmondsangel_hair_pastaacini_de_pepeacai_berryapple_jacks_cerealavocadoaero_chocolate_baragar_agaraubergineangel_food_cakealoo_gobiasiago_cheesebeefbeansbroccolibeetsbarleybaconbiscuitsbundt_cakebreadbagelbutterburritobananabolognablackberryblueberryboysenberrybiscuitbunbarbecuebbqbrowniebuffalo_wingsbasilbok_choybakers_chocolatebreasts_of_chickenbrussel_sproutsblueberry_muffinsbanana_breadbean_saladbaklavabuffalo_meatblack_bean_bunsbalsamic_dressingbreadfruitbirds_nest_cookiesbirds_nest_soupbrainsbean_sproutsblack_eyed_peasbruschettablack_forest_cakeborschtbasmati_ricebulgurbutter_squashbutter_chickencorncarrotscabbagecauliflowercrabcrackercookiescucumbercroutonscantaloupeclamscranberry_saucecorndogcakecashewscoconutcoffeecasserolecrepecodcumquatcaramelcandycurrycocoacream_of_mushroomcroissantscaflibchamomile_teacorn_chipscream_cheesecranberriescurrantscrawfishcrabcakescrème_brulecabbage_rollscalifornia_dressingcassavacreamsiclecoronetocandy_canecoffee_crispcorn_flakescream_puffcows_tonguecaviarcorn_fritterscornbreadcouscouscobblercarrot_cakecustardcompotecoleslawcannolicerealcidercelerycinnamoncaesar_saladcilantrochipscheesecherrychickenchinese_foodchili_con_carnecheeseburgercheetoscheerioschunky_peanut_butterchocolatechowderchow_mien_noodlescheddar_cheesecheesesteakchick_peaschocolate_chipscheesecakechop_sueychapatischutneychurroschalupasduckdonutdanishdrumsticksdutch_apple_pieduriandairydeviled_eggsdilldill_picklesdutch_cookiesdark_chocolatedatesdim_sumdahl_soupdream_whipdolmadesdandelion_greensding_dongsdoritoseggseggplantendiveeggrolleskimo_pieescaroleeclairsescargotespressoenglish_muffinechinaceaelbow_macaronieggnogenchiladaselephant_earseggs_benedictedamamefishfrench_friesfruitfunnel_cakefajitasfortune_cookiefrench_toastfigsfudgeflapjacksfrijolesfrittersfrostingfruitcakefruit_piefrankfurtersfried_piefig_newtonsfritosfettuccinefalafalfree_range_eggsfetafiddlestick_greensflourfricasseeflap_jacksflax_seedsflax_oilfocaccia_breadfudgiclefasolada_soupfried_ricefruit_saladgelatingibletsgingerginger_aleginger_beergingerbreadgelatogarlicgravygreen_beansgrapesgooseberrygarbonzogranolagrapefruitguacamolegrapeseed_oilgraham_crackersgreen_olivesgailangreen_peppergarlic_breadgoats_milkgyozagala_applesguavagumbogazpachogizzardsgoulashhamhorseradishhamburgerhot_doghoagiehoneydewhalibuthuckleberrieshoneyhot_chocolateheinz_ketchuphollandaisehummusherbsherringshearts_of_palmhominyhash_brownshaggishog_jowlhoney_bunho_hosice_creamiced_teaiceeicingindian_jujubeeiceberg_lettuceitalian_breadicing_sugarindian_curryirish_stewimitation_baco_bitsimitation_crabimpastata_cheeseindian_puddinginfant_formulainstant_coffeeinstant_oatmealitalian_dressingjuicejell_ojellyjamjerkyjalapeñojujubesjambalayajamaican_pattieskalekohlrabikiwikidney_beanskumquatkielbasaketchupkamutlima_beanleekslentilsliverlettucelasagnalemonslimeslobsterlinguinelemonadelamblemon_tartslong_john_donutslegumesliverwurstlicoricemustardmacaronimushroommilkmelonmuffinmarshmallowsmozzarellamangomince_meatmarmalademahi_mahimaraschino_cherriesminestronemargarinemaple_syrupmayonnaisemashed_potatoesmackerelmisomanicottimartimonial_cakemoussemars_barmeatloafmasalamilletmoussakamayhawsmarmitenoodlesnutsnachosnectarinesnuggetsnerds_candyneapolitan_ice_creamnalleys_chipsnaan_breadnecco_wafersneufatual_cheeseonionokraorangeoatmealomeletolivesoystersoctopusoreganoorganic_saladox_tailostrich_eggsoat_cakesorzopastapicklespeaspotatopotato_chipsparsleypumpkinpeppersparsnipsporkpopcornpistachiospiepeanutspizzapeanut_butterpuddingpeachespearsplumprunespancakepastrypineapplepeppermintpumpernickelpapayapretzelspopsiclespecanspepperonipersimmonsportabello_mushroomspine_nutsperogiespimentospinto_beanspaprikaputanescaparmesanprawnspoipotstickerspoutinepuffed_wheatprotein_shakesprotein_barspakorapassionfruitporridgeplantainpilafperchpollockpotroastplum_saucepestopeking_duckpaellapita_breadpolentapatepavlovapancettapomegranatequichequailquincequesadillaquinoaquaker_oatmealquick_oatsquakes_rice_cakesquornrutabagarhubarbradishriceribsrollsrye_breadraspberryraviolirigatoniraisinsratatouillerelishritz_crackersrosehipsromainerooibosrice_crackersrotired_deliciousrye_crisprice_paperrootbeer_floatred_lentilsrefried_beansrissottoreuben_sandwichramenradicchioricotta_cheeseromano_cheeseroll_upsspinachsquashsoybeansteakspaghettisaladswiss_cheesesoupsandwichsausagespamscallopsushisaucesuccotashsundaestuffingstrawberrystrombolisnow_peassalamistewsalmonsunflower_seedsswordfishsatsumaslushiessquidsesame_seedsswiss_chardsour_creamsardinessalsasatzikisouvlakisashimispeltsesame_oilspanokopitaspinach_dipsasparillaspumonispartan_applesstarfruitsweet_potatoskors_barssmoresstoned_wheat_crackerssoysaucesmoothie_shakesamosaspanish_ricesnappersplit_peassconesstirfrysorbetsloppy_joesstring_cheesesoufflestarburstshish_kebabshellsshepherds_pieshrimpsherbetshallotshredded_wheatshark_fin_soupshortbreadshortcaketomatoturnipturkeyturtletapiocatoasttacotortillastaffytartstunatangerinetossed_saladtabasco_pepper_saucetamaletv_dinnerstater_totstempuratofutempehtortellinitriscuitstamarindtabbouleh_saladteriyakitortetrifletruffleturmerictahinitwinkiestwizzlerstostitosthighsthousand_island_salad_dressingthimbleberryupside_down_cakeuglifruitunpasteurized_milkvegetablesvenisonvealvermicellividalia_onionvinegarvol_au_ventvegemitevichyssoisevanillavine_leavesvegetable_soupvirgin_oilvakalolovindaloovomit_fruitwatercresswax_beanwatermelonwheat_breadwingswaffleswalnutswaferswalleyewaterwontonswonder_breadwasabiwild_ricewaldorf_saladxmas_cakexiguachex_mixyamyellow_squashyogurtyolkyellow_beansyorkshire_puddingyuccayautiayaki_sobayiyanteszucchinizitizwieback_breadzabaglione_saucezander_fishzoursLastNoun"

var _Noun_index = [...]uint16{0, 9, 14, 21, 30, 40, 56, 66, 81, 88, 104, 117, 127, 145, 152, 170, 179, 188, 203, 212, 225, 229, 234, 242, 247, 253, 258, 266, 276, 281, 286, 292, 299, 305, 312, 322, 331, 342, 349, 352, 360, 363, 370, 383, 388, 396, 412, 430, 445, 462, 474, 484, 491, 503, 518, 535, 545, 563, 578, 584, 596, 611, 621, 638, 645, 657, 663, 676, 690, 694, 701, 708, 719, 723, 730, 737, 745, 753, 763, 768, 783, 790, 794, 801, 808, 814, 823, 828, 831, 838, 845, 850, 855, 860, 877, 887, 893, 906, 916, 928, 939, 947, 955, 964, 976, 989, 1008, 1015, 1025, 1033, 1043, 1055, 1066, 1076, 1087, 1093, 1106, 1115, 1123, 1130, 1141, 1148, 1155, 1163, 1170, 1176, 1181, 1187, 1195, 1207, 1215, 1220, 1226, 1232, 1239, 1251, 1266, 1278, 1285, 1293, 1313, 1322, 1329, 1346, 1360, 1371, 1381, 1396, 1406, 1415, 1423, 1430, 1437, 1445, 1449, 1454, 1460, 1470, 1485, 1491, 1496, 1508, 1512, 1524, 1537, 1551, 1556, 1563, 1572, 1582, 1590, 1606, 1616, 1623, 1627, 1635, 1641, 1648, 1658, 1666, 1673, 1681, 1689, 1703, 1712, 1726, 1732, 1742, 1755, 1768, 1775, 1779, 1791, 1796, 1807, 1814, 1828, 1840, 1844, 1849, 1858, 1866, 1874, 1882, 1891, 1900, 1912, 1921, 1932, 1938, 1948, 1955, 1970, 1974, 1992, 1997, 2006, 2016, 2026, 2034, 2048, 2056, 2069, 2079, 2090, 2097, 2104, 2110, 2120, 2131, 2142, 2148, 2154, 2159, 2170, 2176, 2186, 2194, 2201, 2211, 2220, 2233, 2248, 2260, 2266, 2278, 2290, 2300, 2305, 2316, 2321, 2326, 2334, 2342, 2349, 2352, 2363, 2372, 2379, 2385, 2393, 2400, 2413, 2418, 2431, 2444, 2455, 2461, 2466, 2474, 2488, 2494, 2505, 2511, 2519, 2528, 2534, 2543, 2551, 2555, 2560, 2574, 2589, 2602, 2613, 2625, 2635, 2654, 2668, 2684, 2698, 2712, 2726, 2741, 2757, 2762, 2768, 2773, 2776, 2781, 2790, 2797, 2806, 2822, 2826, 2834, 2838, 2850, 2857, 2865, 2872, 2877, 2886, 2891, 2898, 2903, 2910, 2917, 2923, 2928, 2935, 2943, 2951, 2955, 2966, 2982, 2989, 2999, 3007, 3014, 3022, 3030, 3034, 3039, 3045, 3057, 3067, 3072, 3082, 3091, 3100, 3119, 3129, 3138, 3149, 3159, 3174, 3182, 3186, 3195, 3211, 3217, 3225, 3233, 3239, 3245, 3253, 3260, 3267, 3274, 3278, 3284, 3294, 3301, 3312, 3332, 3345, 3355, 3367, 3383, 3388, 3392, 3398, 3405, 3411, 3417, 3424, 3431, 3438, 3451, 3458, 3470, 3479, 3483, 3488, 3495, 3499, 3505, 3517, 3524, 3531, 3538, 3546, 3550, 3557, 3567, 3570, 3577, 3582, 3595, 3602, 3609, 3614, 3618, 3624, 3631, 3637, 3646, 3656, 3668, 3674, 3682, 3691, 3697, 3706, 3716, 3736, 3745, 3753, 3761, 3772, 3779, 3788, 3796, 3802, 3805, 3816, 3823, 3835, 3849, 3861, 3867, 3879, 3887, 3895, 3900, 3905, 3912, 3920, 3930, 3935, 3946, 3952, 3962, 3969, 3973, 3980, 3988, 3999, 4005, 4010, 4016, 4026, 4032, 4046, 4056, 4073, 4078, 4086, 4093, 4099, 4103, 4107, 4112, 4121, 4130, 4137, 4145, 4152, 4163, 4169, 4182, 4190, 4197, 4204, 4217, 4221, 4234, 4243, 4253, 4267, 4278, 4291, 4299, 4314, 4319, 4328, 4342, 4355, 4363, 4370, 4376, 4383, 4388, 4397, 4402, 4414, 4418, 4426, 4433, 4437, 4444, 4449, 4454, 4463, 4469, 4477, 4487, 4496, 4505, 4511, 4515, 4521, 4536, 4545, 4552, 4560, 4565, 4577, 4588, 4598, 4606, 4611, 4618, 4626, 4633, 4638, 4648, 4659, 4670, 4680, 4687, 4701, 4710, 4722, 4732, 4738, 4759, 4767, 4781, 4787, 4799, 4806, 4816, 4822, 4829, 4835, 4846, 4859, 4866, 4875, 4886, 4892, 4905, 4911, 4918, 4925, 4939, 4953, 4963, 4972, 4978, 4984, 4990, 4996, 5003, 5008, 5012, 5021, 5026, 5031, 5035, 5044, 5056, 5076, 5082, 5092, 5102, 5109, 5113, 5119, 5129, 5138, 5146, 5161, 5169, 5174, 5180, 5187, 5195, 5201, 5209, 5218, 5226, 5232, 5262, 5274, 5290, 5299, 5317, 5327, 5334, 5338, 5348, 5361, 5368, 5379, 5387, 5398, 5405, 5416, 5430, 5440, 5448, 5456, 5467, 5477, 5485, 5495, 5506, 5511, 5518, 5525, 5531, 5538, 5543, 5550, 5562, 5568, 5577, 5590, 5599, 5604, 5612, 5615, 5628, 5634, 5638, 5650, 5667, 5672, 5678, 5687, 5695, 5703, 5707, 5721, 5737, 5748, 5753, 5761}

func (i Noun) String() string {
	if i < 0 || i >= Noun(len(_Noun_index)-1) {
		return fmt.Sprintf("Noun(%d)", i)
	}
	return _Noun_name[_Noun_index[i]:_Noun_index[i+1]]
}