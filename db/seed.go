package db

import (
	"github.com/scmo/apayment-backend/services"
	"github.com/scmo/apayment-backend/models"
)

func Seed_Contributions() {
	//Only seed if table is empty
	if cnt, _ := services.CountContributions(); cnt > 0 {
		return
	}
	services.CreateContribution(&models.Contribution{Code:5416, Name:"Beitrag für besonders Tierfreundliche Stallhaltungssysteme"})
	services.CreateContribution(&models.Contribution{Code:5417, Name:"Beitrag für regelmässigen Auslauf im Freien"})
}

func Seed_ControlPoints() {

	//Only seed if table is empty
	if cnt, _ := services.CountControlCategories(); cnt > 0 {
		return
	}
	if cnt, _ := services.CountPointGroups(); cnt > 0 {
		return
	}
	if cnt, _ := services.CountControlPoints(); cnt > 0 {
		return
	}
	Seed_BTS_Rindergattung_Wasserbueffel()
	Seed_RAUS_Weidetiere();
}

func Seed_BTS_Rindergattung_Wasserbueffel() {
	contribution, _ := services.GetContributionByCode(5416)
	cc := models.ControlCategory{ControlCategoryId:"12.01_2017", Abbreviation:"BTS - Rindergattung und Wasserbüffel", Contribution: contribution}
	services.CreateControlCategory(&cc)

	pg1 := models.PointGroup{PointGroupId:"A1", Abbreviation:"Rinder - Milchkühe", PointGroup: "Rinder - Milchkühe", PointGroupCode:1110, ControlCategory: &cc}
	services.CreatePointGroup(&pg1)

	cp1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten",
		ControlPoint:"Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg1}
	services.CreateControlPoint(&cp1)
	cp2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Mind. 15 Lux Tageslicht im Stall",
		ControlPoint:"Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg1}
	services.CreateControlPoint(&cp2)

	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points:60, ControlPoint:&cp1})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cp1})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cp1})
	services.CreateLack(&models.Lack{Name:"Etwas zu wenig Tageslicht", Points:10, ControlPoint:&cp2})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig Tageslicht", Points:110, ControlPoint:&cp2})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cp2})


	// A2	Rinder - andere Kühe
	pg2 := models.PointGroup{PointGroupId:"A2", Abbreviation:"Rinder - andere Kühe", PointGroup: "Rinder - andere Kühe", PointGroupCode:1150, ControlCategory: &cc}
	services.CreatePointGroup(&pg2)

	cpA2_1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten",
		ControlPoint:"Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_1)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA2_1})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA2_1})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA2_1})

	cpA2_2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Mind. 15 Lux Tageslicht im Stall", ControlPoint:"Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_2)
	services.CreateLack(&models.Lack{Name:"Etwas zu wenig Tageslicht", Points:10, ControlPoint:&cpA2_2})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig Tageslicht", Points:110, ControlPoint:&cpA2_2})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA2_2})

	cpA2_3 := models.ControlPoint{ControlPointId:"03", Abbreviation:"Fress- und Tränkebereich: befestigter Boden", ControlPoint:"Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_3)
	services.CreateLack(&models.Lack{Name:"Fress- und Tränkebereich: unbefestigter Boden", Points:110, ControlPoint:&cpA2_3})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA2_3})

	cpA2_4 := models.ControlPoint{ControlPointId:"04", Abbreviation:"Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint:"Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_4)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA2_4})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA2_4})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA2_4})

	cpA2_5 := models.ControlPoint{ControlPointId:"05.1", Abbreviation:"Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_5)

	cpA2_6 := models.ControlPoint{ControlPointId:"05.1.1", Abbreviation:"Liegemattenfabrikat BTS-konform", ControlPoint:"Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_6)
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points:60, ControlPoint:&cpA2_6})
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points:110, ControlPoint:&cpA2_6})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA2_6})

	cpA2_7 := models.ControlPoint{ControlPointId:"05.1.2", Abbreviation:"Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_7)
	services.CreateLack(&models.Lack{Name:"Zu wenig BTS-konforme Einstreu", Points:10, ControlPoint:&cpA2_7})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig BTS-konforme Einstreu", Points:60, ControlPoint:&cpA2_7})
	services.CreateLack(&models.Lack{Name:"Keine BTS-konforme Einstreu", Points:110, ControlPoint:&cpA2_7})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA2_7})

	cpA2_8 := models.ControlPoint{ControlPointId:"05.2", Abbreviation:"Liegebereich in allen anderen Laufställen	", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_8)

	cpA2_9 := models.ControlPoint{ControlPointId:"05.2.1", Abbreviation:"Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint:"Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_9)
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points:60, ControlPoint:&cpA2_9})
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points:110, ControlPoint:&cpA2_9})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA2_9})

	/*
		A3 - Rinder - weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung
	 */
	pg3 := models.PointGroup{PointGroupId:"A3",
		Abbreviation:"Rinder - weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung",
		PointGroup: "weibliche Tiere über 365 - 730 Tage alt, ohne Abkalbung", PointGroupCode:1128, ControlCategory: &cc}
	services.CreatePointGroup(&pg3)

	cpA3_1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten", ControlPoint:"Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_1)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA3_1})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA3_1})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA3_1})
	cpA3_2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Mind. 15 Lux Tageslicht im Stall", ControlPoint:"Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_2)
	services.CreateLack(&models.Lack{Name:"Etwas zu wenig Tageslicht", Points:10, ControlPoint:&cpA3_2})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig Tageslicht", Points:110, ControlPoint:&cpA3_2})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA3_2})
	cpA3_3 := models.ControlPoint{ControlPointId:"03", Abbreviation:"Fress- und Tränkebereich: befestigter Boden", ControlPoint:"Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_3)
	services.CreateLack(&models.Lack{Name:"Fress- und Tränkebereich: unbefestigter Boden", Points:110, ControlPoint:&cpA3_3})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA3_3})
	cpA3_4 := models.ControlPoint{ControlPointId:"04", Abbreviation:"Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint:"Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_4)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA3_4})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA3_4})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA3_4})
	cpA3_5 := models.ControlPoint{ControlPointId:"05.1", Abbreviation:"Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_5)
	cpA3_6 := models.ControlPoint{ControlPointId:"05.1.1", Abbreviation:"Liegemattenfabrikat BTS-konform", ControlPoint:"Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_6)
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points:60, ControlPoint:&cpA3_6})
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points:110, ControlPoint:&cpA3_6})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA3_6})
	cpA3_7 := models.ControlPoint{ControlPointId:"05.1.2", Abbreviation:"Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_7)
	services.CreateLack(&models.Lack{Name:"Zu wenig BTS-konforme Einstreu", Points:10, ControlPoint:&cpA3_7})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig BTS-konforme Einstreu", Points:60, ControlPoint:&cpA3_7})
	services.CreateLack(&models.Lack{Name:"Keine BTS-konforme Einstreu", Points:110, ControlPoint:&cpA3_7})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA3_7})
	cpA3_8 := models.ControlPoint{ControlPointId:"05.2", Abbreviation:"Liegebereich in allen anderen Laufställen	", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_8)
	cpA3_9 := models.ControlPoint{ControlPointId:"05.2.1", Abbreviation:"Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint:"Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_9)
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points:60, ControlPoint:&cpA3_9})
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points:110, ControlPoint:&cpA3_9})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA3_9})



	/*
	 	A4 = Rinder - weibliche Tiere, über 160-365 Tage alt
	  */
	pg4 := models.PointGroup{PointGroupId:"A4", Abbreviation:"Rinder - weibliche Tiere, über 160-365 Tage alt", PointGroup: "weibliche Tiere über 160 - 365 Tage alt", PointGroupCode:1141, ControlCategory: &cc}
	services.CreatePointGroup(&pg4)

	cpA4_1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten", ControlPoint:"Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_1)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA4_1})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA4_1})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA4_1})
	cpA4_2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Mind. 15 Lux Tageslicht im Stall", ControlPoint:"Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_2)
	services.CreateLack(&models.Lack{Name:"Etwas zu wenig Tageslicht", Points:10, ControlPoint:&cpA4_2})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig Tageslicht", Points:110, ControlPoint:&cpA4_2})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA4_2})
	cpA4_3 := models.ControlPoint{ControlPointId:"03", Abbreviation:"Fress- und Tränkebereich: befestigter Boden", ControlPoint:"Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_3)
	services.CreateLack(&models.Lack{Name:"Fress- und Tränkebereich: unbefestigter Boden", Points:110, ControlPoint:&cpA4_3})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA4_3})
	cpA4_4 := models.ControlPoint{ControlPointId:"04", Abbreviation:"Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint:"Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_4)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA4_4})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA4_4})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA4_4})
	cpA4_5 := models.ControlPoint{ControlPointId:"05.1", Abbreviation:"Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_5)
	cpA4_6 := models.ControlPoint{ControlPointId:"05.1.1", Abbreviation:"Liegemattenfabrikat BTS-konform", ControlPoint:"Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_6)
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points:60, ControlPoint:&cpA4_6})
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points:110, ControlPoint:&cpA4_6})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA4_6})
	cpA4_7 := models.ControlPoint{ControlPointId:"05.1.2", Abbreviation:"Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_7)
	services.CreateLack(&models.Lack{Name:"Zu wenig BTS-konforme Einstreu", Points:10, ControlPoint:&cpA4_7})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig BTS-konforme Einstreu", Points:60, ControlPoint:&cpA4_7})
	services.CreateLack(&models.Lack{Name:"Keine BTS-konforme Einstreu", Points:110, ControlPoint:&cpA4_7})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA4_7})
	cpA4_8 := models.ControlPoint{ControlPointId:"05.2", Abbreviation:"Liegebereich in allen anderen Laufställen	", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_8)
	cpA4_9 := models.ControlPoint{ControlPointId:"05.2.1", Abbreviation:"Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint:"Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_9)
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points:60, ControlPoint:&cpA4_9})
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points:110, ControlPoint:&cpA4_9})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA4_9})

	/*
		A6 - Rinder - männliche Tiere, über 730 Tage alt
	 */
	pg6 := models.PointGroup{PointGroupId:"A6", Abbreviation:"Rinder - männliche Tiere, über 730 Tage alt", PointGroup: "männliche Tiere, über 730 Tage alt", PointGroupCode:1124, ControlCategory: &cc}
	services.CreatePointGroup(&pg6)

	cpA6_1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten", ControlPoint:"Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_1)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA6_1})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA6_1})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA6_1})
	cpA6_2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Mind. 15 Lux Tageslicht im Stall", ControlPoint:"Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_2)
	services.CreateLack(&models.Lack{Name:"Etwas zu wenig Tageslicht", Points:10, ControlPoint:&cpA6_2})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig Tageslicht", Points:110, ControlPoint:&cpA6_2})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA6_2})
	cpA6_3 := models.ControlPoint{ControlPointId:"03", Abbreviation:"Fress- und Tränkebereich: befestigter Boden", ControlPoint:"Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_3)
	services.CreateLack(&models.Lack{Name:"Fress- und Tränkebereich: unbefestigter Boden", Points:110, ControlPoint:&cpA6_3})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA6_3})
	cpA6_4 := models.ControlPoint{ControlPointId:"04", Abbreviation:"Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint:"Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_4)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA6_4})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA6_4})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA6_4})
	cpA6_5 := models.ControlPoint{ControlPointId:"05.1", Abbreviation:"Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_5)
	cpA6_6 := models.ControlPoint{ControlPointId:"05.1.1", Abbreviation:"Liegemattenfabrikat BTS-konform", ControlPoint:"Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_6)
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points:60, ControlPoint:&cpA6_6})
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points:110, ControlPoint:&cpA6_6})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA6_6})
	cpA6_7 := models.ControlPoint{ControlPointId:"05.1.2", Abbreviation:"Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_7)
	services.CreateLack(&models.Lack{Name:"Zu wenig BTS-konforme Einstreu", Points:10, ControlPoint:&cpA6_7})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig BTS-konforme Einstreu", Points:60, ControlPoint:&cpA6_7})
	services.CreateLack(&models.Lack{Name:"Keine BTS-konforme Einstreu", Points:110, ControlPoint:&cpA6_7})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA6_7})
	cpA6_8 := models.ControlPoint{ControlPointId:"05.2", Abbreviation:"Liegebereich in allen anderen Laufställen	", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_8)
	cpA6_9 := models.ControlPoint{ControlPointId:"05.2.1", Abbreviation:"Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint:"Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_9)
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points:60, ControlPoint:&cpA6_9})
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points:110, ControlPoint:&cpA6_9})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA6_9})

	/*
		A7 - Rinder - männliche Tiere, über 365- 730 Tage alt
	 */
	pg7 := models.PointGroup{PointGroupId:"A7", Abbreviation:"Rinder - männliche Tiere, über 365- 730 Tage alt", PointGroup: "männliche Tiere, über 365 bis 730 Tage alt", PointGroupCode:1129, ControlCategory: &cc}
	services.CreatePointGroup(&pg7)

	cpA7_1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten", ControlPoint:"Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_1)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA7_1})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA7_1})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA7_1})
	cpA7_2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Mind. 15 Lux Tageslicht im Stall", ControlPoint:"Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_2)
	services.CreateLack(&models.Lack{Name:"Etwas zu wenig Tageslicht", Points:10, ControlPoint:&cpA7_2})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig Tageslicht", Points:110, ControlPoint:&cpA7_2})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA7_2})
	cpA7_3 := models.ControlPoint{ControlPointId:"03", Abbreviation:"Fress- und Tränkebereich: befestigter Boden", ControlPoint:"Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_3)
	services.CreateLack(&models.Lack{Name:"Fress- und Tränkebereich: unbefestigter Boden", Points:110, ControlPoint:&cpA7_3})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA7_3})
	cpA7_4 := models.ControlPoint{ControlPointId:"04", Abbreviation:"Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint:"Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_4)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA7_4})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA7_4})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA7_4})
	cpA7_5 := models.ControlPoint{ControlPointId:"05.1", Abbreviation:"Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_5)
	cpA7_6 := models.ControlPoint{ControlPointId:"05.1.1", Abbreviation:"Liegemattenfabrikat BTS-konform", ControlPoint:"Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_6)
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points:60, ControlPoint:&cpA7_6})
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points:110, ControlPoint:&cpA7_6})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA7_6})
	cpA7_7 := models.ControlPoint{ControlPointId:"05.1.2", Abbreviation:"Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_7)
	services.CreateLack(&models.Lack{Name:"Zu wenig BTS-konforme Einstreu", Points:10, ControlPoint:&cpA7_7})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig BTS-konforme Einstreu", Points:60, ControlPoint:&cpA7_7})
	services.CreateLack(&models.Lack{Name:"Keine BTS-konforme Einstreu", Points:110, ControlPoint:&cpA7_7})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA7_7})
	cpA7_8 := models.ControlPoint{ControlPointId:"05.2", Abbreviation:"Liegebereich in allen anderen Laufställen	", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_8)
	cpA7_9 := models.ControlPoint{ControlPointId:"05.2.1", Abbreviation:"Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint:"Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_9)
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points:60, ControlPoint:&cpA7_9})
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points:110, ControlPoint:&cpA7_9})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA7_9})
	/*
		A8 - Rinder - männliche Tiere, über 160-365 Tage alt
	 */
	pg8 := models.PointGroup{PointGroupId:"A8", Abbreviation:"Rinder - männliche Tiere, über 160-365 Tage alt", PointGroup: "männliche Tiere, über 160 bis 365 Tage alt", PointGroupCode:1143, ControlCategory: &cc}
	services.CreatePointGroup(&pg8)
	cpA8_1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten", ControlPoint:"Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_1)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA8_1})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA8_1})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA8_1})
	cpA8_2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Mind. 15 Lux Tageslicht im Stall", ControlPoint:"Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_2)
	services.CreateLack(&models.Lack{Name:"Etwas zu wenig Tageslicht", Points:10, ControlPoint:&cpA8_2})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig Tageslicht", Points:110, ControlPoint:&cpA8_2})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA8_2})
	cpA8_3 := models.ControlPoint{ControlPointId:"03", Abbreviation:"Fress- und Tränkebereich: befestigter Boden", ControlPoint:"Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_3)
	services.CreateLack(&models.Lack{Name:"Fress- und Tränkebereich: unbefestigter Boden", Points:110, ControlPoint:&cpA8_3})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA8_3})
	cpA8_4 := models.ControlPoint{ControlPointId:"04", Abbreviation:"Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint:"Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_4)
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points:60, ControlPoint:&cpA8_4})
	services.CreateLack(&models.Lack{Name:"Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points:110, ControlPoint:&cpA8_4})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA8_4})
	cpA8_5 := models.ControlPoint{ControlPointId:"05.1", Abbreviation:"Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_5)
	cpA8_6 := models.ControlPoint{ControlPointId:"05.1.1", Abbreviation:"Liegemattenfabrikat BTS-konform", ControlPoint:"Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_6)
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points:60, ControlPoint:&cpA8_6})
	services.CreateLack(&models.Lack{Name:"Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points:110, ControlPoint:&cpA8_6})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA8_6})
	cpA8_7 := models.ControlPoint{ControlPointId:"05.1.2", Abbreviation:"Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_7)
	services.CreateLack(&models.Lack{Name:"Zu wenig BTS-konforme Einstreu", Points:10, ControlPoint:&cpA8_7})
	services.CreateLack(&models.Lack{Name:"Viel zu wenig BTS-konforme Einstreu", Points:60, ControlPoint:&cpA8_7})
	services.CreateLack(&models.Lack{Name:"Keine BTS-konforme Einstreu", Points:110, ControlPoint:&cpA8_7})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA8_7})
	cpA8_8 := models.ControlPoint{ControlPointId:"05.2", Abbreviation:"Liegebereich in allen anderen Laufställen	", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_8)
	cpA8_9 := models.ControlPoint{ControlPointId:"05.2.1", Abbreviation:"Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint:"Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_9)
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points:60, ControlPoint:&cpA8_9})
	services.CreateLack(&models.Lack{Name:"Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points:110, ControlPoint:&cpA8_9})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cpA8_9})
}

func Seed_RAUS_Weidetiere() {
	contribution, _ := services.GetContributionByCode(5417)
	cc := models.ControlCategory{ControlCategoryId:"12.07_2017", Abbreviation:"RAUS-Weidetiere", Contribution: contribution}
	services.CreateControlCategory(&cc)

	pg1 := models.PointGroup{PointGroupId:"A1", Abbreviation:"Milchkühe", PointGroupCode:1110, ControlCategory: &cc}
	services.CreatePointGroup(&pg1)

	cp1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten",
		ControlPoint:"Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen). Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen). Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen). * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg1}
	services.CreateControlPoint(&cp1)
	cp2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen",
		ControlPoint:"Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen)Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen)Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen)* = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg1}
	services.CreateControlPoint(&cp2)

	cp3 := models.ControlPoint{ControlPointId:"03", Abbreviation:"Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und  - entsprechen den Anforderungen",
		ControlPoint:"Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen)Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen)Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen)Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt seinKategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg1}
	services.CreateControlPoint(&cp3)

	services.CreateLack(&models.Lack{Name:"Laufhof befindet sich nicht im Freien", Points:110, ControlPoint:&cp1})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cp1})

	services.CreateLack(&models.Lack{Name:"Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points:60, ControlPoint:&cp2})
	services.CreateLack(&models.Lack{Name:"Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points:110, ControlPoint:&cp2})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cp2})

	services.CreateLack(&models.Lack{Name:"Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points:60, ControlPoint:&cp3})
	services.CreateLack(&models.Lack{Name:"Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points:110, ControlPoint:&cp3})
	services.CreateLack(&models.Lack{Name:"Anderer Mangel", ControlPoint:&cp3})




	// A3 - Rinder - weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung
	pg3 := models.PointGroup{PointGroupId:"A3", Abbreviation:"Rinder - weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung", PointGroup: "weibliche Tiere über 365 - 730 Tage alt, ohne Abkalbung", PointGroupCode:1128, ControlCategory: &cc}
	services.CreatePointGroup(&pg3)

	// A4 = Rinder - weibliche Tiere, über 160-365 Tage alt
	pg4 := models.PointGroup{PointGroupId:"A4", Abbreviation:"Rinder - weibliche Tiere, über 160-365 Tage alt", PointGroup: "weibliche Tiere über 160 - 365 Tage alt", PointGroupCode:1141, ControlCategory: &cc}
	services.CreatePointGroup(&pg4)

	// A5 - Rinder - weibliche Tiere, bis 160 Tage alt
	pg5 := models.PointGroup{PointGroupId:"A5", Abbreviation:"Rinder - weibliche Tiere, bis 160 Tage alt", PointGroup: "weibliche Tiere bis 160 Tage alt", PointGroupCode:1142, ControlCategory: &cc}
	services.CreatePointGroup(&pg5)

	// A6 - Rinder - männliche Tiere, über 730 Tage alt
	pg6 := models.PointGroup{PointGroupId:"A6", Abbreviation:"Rinder - männliche Tiere, über 730 Tage alt", PointGroup: "männliche Tiere, über 730 Tage alt", PointGroupCode:1124, ControlCategory: &cc}
	services.CreatePointGroup(&pg6)

	// A7 - Rinder - männliche Tiere, über 365- 730 Tage alt
	pg7 := models.PointGroup{PointGroupId:"A7", Abbreviation:"Rinder - männliche Tiere, über 365- 730 Tage alt", PointGroup: "männliche Tiere, über 365 bis 730 Tage alt", PointGroupCode:1129, ControlCategory: &cc}
	services.CreatePointGroup(&pg7)

	// A8 - Rinder - männliche Tiere, über 160-365 Tage alt
	pg8 := models.PointGroup{PointGroupId:"A8", Abbreviation:"Rinder - männliche Tiere, über 160-365 Tage alt", PointGroup: "männliche Tiere, über 160 bis 365 Tage alt", PointGroupCode:1143, ControlCategory: &cc}
	services.CreatePointGroup(&pg8)

	// A9 - Rinder - männliche Tiere, bis 160 Tage alt
	pg9 := models.PointGroup{PointGroupId:"A8", Abbreviation:"Rinder - männliche Tiere, bis 160 Tage alt", PointGroup: "männliche Tiere, bis 160 Tage alt", PointGroupCode:1144, ControlCategory: &cc}
	services.CreatePointGroup(&pg9)
}

func Seed_Users() {
	if cnt, _ := services.CountRoles(); cnt > 0 {
		return
	}
	admin := models.Role{Name:"Admin"}
	farmer := models.Role{Name:"Farmer"}
	canton := models.Role{Name:"Canton"}
	inspector := models.Role{Name:"Inspector"}

	services.CreateRole(&admin);
	services.CreateRole(&farmer);
	services.CreateRole(&canton);
	services.CreateRole(&inspector);

	if cnt, _ := services.CountUsers(); cnt > 0 {
		return
	}

	roles := make([]*models.Role, 1)
	roles[0] = &farmer
	services.CreateUser(&models.User{Username:"farmer1", Password:"farmer1", Firstname: "Florian", Lastname:"Meisterhans", TVD: 1015010, Roles: roles })
	services.CreateUser(&models.User{Username:"farmer2", Password:"farmer2", Firstname: "Max", Lastname:"Keller", TVD: 1015010, Roles: roles})

	roles = make([]*models.Role, 1)
	roles[0] = &inspector
	services.CreateUser(&models.User{Username:"inspect", Password:"inspect", Firstname: "Inspector", Lastname: "Gadget", Roles: roles})

	roles = make([]*models.Role, 1)
	roles[0] = &admin
	services.CreateUser(&models.User{Username:"admin", Password:"admin", Firstname: "Admin", Lastname: "Admin", Roles: roles})

	roles = make([]*models.Role, 1)
	roles[0] = &canton
	services.CreateUser(&models.User{Username:"canton", Password:"canton", Firstname: "Maria", Lastname: "Walder", Roles: roles})

}

func Seed_LegalForm() {
	if cnt, _ := services.CountLegalForms(); cnt > 0 {
		return
	}
	services.CreateLegalForm(&models.LegalForm{Code: 1, Name:"Natürtliche Person"})
	services.CreateLegalForm(&models.LegalForm{Code: 2, Name:"Einfache Gesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 3, Name:"Kollektivgesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 4, Name:"Kommanditgesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 5, Name:"Kommanditaktiengesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 6, Name:"Aktiengesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 7, Name:"Gesellschaft mit beschränkter Haftung"})
	services.CreateLegalForm(&models.LegalForm{Code: 8, Name:"Genossenschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 9, Name:"Verein, Vereinigung"})
	services.CreateLegalForm(&models.LegalForm{Code: 10, Name:"Stiftung"})

	services.CreateLegalForm(&models.LegalForm{Code: 24, Name:"Öffentlich-rechtliche Körperschaft(Verwaltung)"})
	services.CreateLegalForm(&models.LegalForm{Code: 25, Name:"Staatlich anerkannte Landeskirche"})
	services.CreateLegalForm(&models.LegalForm{Code: 30, Name:"Bund (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 31, Name:"Kanton (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 32, Name:"Bezirk (Betrieb)"})

	services.CreateLegalForm(&models.LegalForm{Code: 33, Name:"Gemeinde (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 34, Name:"Öffentlich-rechtliche Körperschaft (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 99, Name:"Nicht zugeteilt"})
}

func Seed_PlantType() {
	if cnt, _ := services.CountPlantTypes(); cnt > 0 {
		return
	}
	services.CreatePlantType(&models.PlantType{Code: 1, Name:"Ganzjahresbetrieb"})
	services.CreatePlantType(&models.PlantType{Code: 2, Name:"Produktsstätte"})
	services.CreatePlantType(&models.PlantType{Code: 3, Name:"Betriebsgemeinschaft"})
}
