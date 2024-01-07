package dataset

import (
	"event-hunters/models"
	"time"

	"github.com/volatiletech/null/v8"
)

func InitializeEvents() []models.Event {
	EventsToInsert := []models.Event{
		{
			EventcreatorID:  null.NewInt(1, true),
			Location:        null.NewString("Ciputra Artpreneur, Ciputra World 1, Lt. 10 Retail Podium, Jl. Professor Doktor Satrio Kav. 3-5, Setiabudi, Jakarta Selatan, Jakarta, Indonesia, Theater Ciputra Artpreneur", true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 17, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 25, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 9, 11, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 1, 30, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("Konser ELFAS SINGERS Gempita Cinta", true),
			Description:     null.NewString("Menandai 45 tahun eksistensi di dunia musik Tanah Air, ELFA'S SINGERS bakal menggelar Konser Musik yang diberi tajuk : GEMPITA CINTA. Panggung musik istimewa ini akan digelar Rabu, 20 Desember 2023 di Theater Ciputra Artpreneur, Kuningan, Jakarta Selatan, pada pukul 19.30 - 22.00 WIB. Penonton diajak mengarungi waktu dari masa ke masa bersama ELFA'S SINGERS, tak pelak suguhan konser kolaborasi apik dengan Dian HP ini akan menjadi panggung musik yang istimewa. ELFA'S SINGERS yang beranggotakan: Yana Julio, Lita Zen, Agus Wisman dan Ucie Nurul, bukan sekedar mengajak penonton bernostalgia dengan lagu-lagu hits mereka di masa-masa silam. menyanyi dengan harmonisasi vokal indah di atas panggung, melekat pada ELFA'S SINGERS. Hal tersebut menjadi inspirasi keren untuk penyanyi generasi milenial atau Gen Z. Yuk nonton sama keluarga / pacar / calon pendamping hidup / teman yang kece, ajak semuanya yah, kalian akan merasakan cinta & kebahagiaan. Sampai ketemu di Artpreneur", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"music", "concert", "eventjakarta", "musik"},
			FeaturedImages:  []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702389367-yAFLAv0mteaSAI9QUFrKfaTUV1Dlu7hr.jpg?width=1024&quality=90", "https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702389402-T7T2674YtR11LJGQVIiKhFg2RpMhTt8m.jpg?width=1024&quality=90", "https://www.goersapp.com/events/konser-elfas-singers-gempita-cinta--konserelfassingers"},
			GuestStar:       []string{"Dian HP", "The SECIORIAS", "MYLENE", "ELFA'S JAZZ & POP SINGERS"},
		},
		{
			EventcreatorID:  null.NewInt(1, true),
			Location:        null.NewString("Universitas Nasional,  Jl. Sawo Manila, Pasar Minggu, Jakarta Selatan, Jakarta, Indonesia Lantai 1", true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 14, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("STAND UP COMEDY LIMA EMPAT ABDEL ACHRIAN", true),
			Description:     null.NewString("\"LIMA EMPAT adalah stand up comedy show pertama cing abdel ( Abdel Achrian ). sebuah cerita perjalanan seorang anak minang yang lahir di pisangan. LIMA EMPAT akan menjadi sebuah pertunjukan perjalanan seorang seniman komedi bernama Abdel Achrian, dari belia hingga sekarang menjelang usia senja. lima empat sebuah angka yang tidak mudah ditempuh begitu saja, melainkan penuh kisah perjuangan, keluh kesah & cucur keringat. sampai jumpa di LIMA EMPAT sebuah angka setelah lima tiga & sebelumlimalima \"", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"standupcomedy", "abdelachrian", "eventjakarta"},
			FeaturedImages:  []string{"https://www.goersapp.com/events/stand-up-comedy-lima-empat-abdel-achrian--limaempatshow", "https://www.goersapp.com/events/stand-up-comedy-lima-empat-abdel-achrian--limaempatshow"},
			GuestStar:       []string{"Abdel Achrian"},
		},
		{
			EventcreatorID:  null.NewInt(1, true),
			Location:        null.NewString("Jakarta, Indonesia,  Setia Budi, Jakarta Selatan Kota, Jakarta, Indonesia", true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 4, 2, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("Growth Hacking Summit 2024", true),
			Description:     null.NewString("Bergabunglah bersama kami di acara premier Indonesia bersama lebih dari 500 profesional pemasaran pertumbuhan, pendiri, dan para ahli produk yang memiliki pemikiran sejenis. Growth Hacking Summit 2024 bukan hanya konferensi; ini adalah perpaduan unik antara inovasi, inspirasi, dan wawasan yang dapat diimplementasikan!", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"edukasi", "seanellis", "innovasi"},
			FeaturedImages:  []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702877844-R1gaXdUdFUAo7MEv3dZcci5Ds8bAm2EX.png?width=1024&quality=90"},
			GuestStar:       []string{"Sean Ellis"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("AYANA Midplaza Jakarta, Kav 10-11 Jalan Jenderal Sudirman Kecamatan Tanah Abang, Daerah Khusus Ibukota Jakarta 10220", true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 3, 2, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 3, 2, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 16, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("Begin Edu Fair Jakarta (Indonesia)", true),
			Description:     null.NewString("Begin Edu Fair - Jakarta adalah platform bagi pelajar dan orang tua mereka, serta mahasiswa universitas yang sedang mencari program studi internasional, termasuk kursus persiapan, sarjana, magister, MBA, serta program bahasa dan musim panas. Jangan lewatkan kesempatan untuk bertemu dengan calon mahasiswa dan orang tua mereka yang tertarik dengan program di institusi Anda. Anda akan menerima daftar kontak dari mereka yang menunjukkan minat pada institusi Anda.", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"edukasi", "edufair", "educationfair", "internationalstudyprogram"},
			FeaturedImages:  []string{"https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F543006069%2F143238831910%2F1%2Foriginal.20230626-143225?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C1536%2C768&s=1d7bd63cadf03dea1bf1de4d8625494f"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("Universitas Nasional,  Jl. Sawo Manila, Pasar Minggu, Jakarta Selatan, Jakarta, Indonesia Lantai 1", true),
			Category:        []string{"Art & Culture", "Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 3, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 17, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("Together on the Wheel - Pottery Class by Classpop!", true),
			Description:     null.NewString("Kelas kerajinan tanah liat yang memuaskan dan menenangkan dipandu oleh seorang instruktur kerajinan tanah liat berpengalaman melalui Classpop! Pelajari berbagai pendekatan dalam kerajinan ini dan ciptakan karya seni unik. Bergabunglah dengan Classpop! untuk \"Bersama di Roda\" dengan Instruktur Michelle dan telusuri teknik-teknik menarik serta gaya seni dalam kelas kerajinan tanah liat yang dirancang untuk pemula maupun pengrajin berpengalaman. Anda akan menciptakan karya kustom yang akan melayani Anda selama bertahun-tahun, semua dengan bantuan instruktur yang ahli dan bersemangat untuk berbagi pengetahuan mereka tentang kerajinan ini dengan siswa seperti Anda. Classpop! menawarkan kelas-kelas kreatif yang penuh kesenangan, seperti tari, lukisan, kerajinan tanah liat, dan lainnya, di lingkungan yang santai dan dipimpin oleh instruktur kelas dunia. Pelajari, ciptakan, dan dapatkan inspirasi dengan memesan kelas hari ini! Untuk informasi tambahan dan tanggal kunjungi Together on the Wheel!", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"pottery", "kerajinan", "tanahliat", "art", "seni"},
			FeaturedImages:  []string{"https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F626529949%2F316402407909%2F1%2Foriginal.20231023-152114?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C2160%2C1080&s=40a43805d992d206c7c72c5d24f0614b"},
		},
		{
			EventcreatorID:  null.NewInt(1, true),
			Location:        null.NewString("Gambir Expo Kemayoran Jl. Benyamin Suaeb, Pademangan Tim., Kec. Pademangan, Jakarta Utara", true),
			Category:        []string{"Entertainment & Performance", "Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 3, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("MOTION IME FESTIVAL", true),
			Description:     null.NewString("Kamu tau gak, kalau banyak kabupaten di Papua masuk dalam daftar daerah tertingal? Hal ini tidak lain karena akses pendidikan yang tidak merata dan Pemberdayaan SDM yang belum optimal! Melalui MOTION IME FEST 2023, berkolaborasi dengan Yayasan Cakra Abhipraya Responsif, memiliki inisiatif untuk membuat sebuah Konser amal, yang seluruh dari hasil penjualan tiketnya nanti akan di donasikan untuk membuat education center, di daerah papua pegunungan, yang di konsep berupa sekolah alam, untuk mendukung peningkatan mutu pendidikan dan pengembangan generasi penerus papua. Program ini akan di jalankan awal tahun 2024 di Distrik walaik, Kabupaten Jayawijaya Provinsi Papua Pegunungan. Selain program pendidikan kami juga akan menjalankan program gizi untuk menurunkan angka stunting. Ayo sebarkan kabar baik ini dan jadi lah bagian dari kebaikan dengan, datang ke acara MOTION IME FEST 2023 2-3 Desember 2023 di Gambir Expo Kemayoran dan nikmati berbagai macam pertunjukan termasuk cosplay, konser dan talkshow", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"cosplay", "hakken", "windahbasudara", "eventcosplay", "charity"},
			FeaturedImages:  []string{"https://ik.imagekit.io/playstagingid/event/banner-1699796688018_0cTNNQadT"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("Hotel Grand Tjokro, Jalan Daan Mogot No.63 Tj. Duren Utara Kec. Grogol petamburan, Daerah Khusus Ibukota Jakarta 11470", true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 3, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 3, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 5, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 19, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("Webinar INTERNET MARKETING REVOLUTION", true),
			Description:     null.NewString("INTERNET MARKETING REVOLUTION Temukan 3 Rahasia SAKTI Bagaimana Memulai Bisnis di Era NEW NORMAL Tanpa Perlu Modal Besar Bahkan Langsung Bisa Anda Praktekkan Setelah Webinar! Anda akan Belajar langsung dari Praktisi Digital Marketing Sejak 2010 dan Sudah Membantu Ribuan Newbie Untuk Memiliki Bisnis Digital TEMUKAN SOLUSINYA dalam Gratis Webinar 90 menit:*  Temukan 3 RAHASIA DALAM WEBINAR ini sbb : RAHASIA #1 Temukan 3 Langkah Sederhana Memiliki Bisnis Ideal di Era New Normal RAHASIA #2 Miliki 4 Persyaratan Terpenting Supaya Bisnis Anda Kebal Krisis RAHASIA #3 Rahasia Top Internet Millionaire Dunia yang Bisa Anda Duplikasi Polanya Karena jumlah Tiket Webinar Sangat Terbatas, Amankan Tiket GRATIS Anda sekarang juga :", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"webinar", "internet", "marketing", "revolution", "bisnis"},
			FeaturedImages:  []string{"https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F293461349%2F312123412152%2F1%2Foriginal.20200605-100422?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C200%2C1080%2C540&s=94fe73bbb2fa3bb25ee5de9350241936"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("SMAN 3 TARUNA ANGKASA JAWA TIMUR Sman3 Taruna Angkasa, Jalan Ring Road Barat, Ngegong, Manguharjo, Mangu Harjo, Madiun Kota, Jawa Timur, Indonesia", true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 5, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 6, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("CLOSING SISO 40", true),
			Description:     null.NewString("Closing SISO 40 adalah acara tahunan dari OSISSMAGANASA yang berfungsi sebagai bentuk perayaan dan puncak peringatan ulang tahun SMAN 3 Taruna Angkasa Jawa Timur. Dalam acara ini, kami akan mengumpulkan penonton untuk menonton pertunjukan dari ekstrakurikuler sekolah dan bintang tamu terakhir kami, yaitu \"FOR REVENGE.\"", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"konser", "musik", "konsermusik"},
			FeaturedImages:  []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1701940501-koPp8pYC0gpuCNOZXaGdKHJu85rwd21v.png?width=1024&quality=90"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("SMA Notre Dame, Jakarta Barat", true),
			Category:        []string{"Sport", "Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Now(), true),
			EndregisterDate: null.NewTime(time.Now().AddDate(0, 0, 7), true),
			StartEvent:      null.NewTime(time.Now().AddDate(0, 0, 10), true),
			EndEvent:        null.NewTime(time.Now().AddDate(0, 0, 12), true),
			CreatedAt:       null.NewTime(time.Now(), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("0.0", false),
			Longitude:       null.NewString("0.0", false),
			Title:           null.NewString("Event Cup Sekolah Notre Dame Querencia 2024", true),
			Description: null.NewString(`Selamat datang di Event Cup Sekolah Notre Dame! Sebuah perhelatan prestisius yang mengundang seluruh komunitas sekolah dan pecinta olahraga untuk merayakan semangat persaingan, kebersamaan, dan prestasi di lingkungan pendidikan.
		
		Dalam atmosfer yang penuh semangat, kami dengan bangga menghadirkan serangkaian perlombaan yang menantang dan memacu kemampuan atletik para peserta. Terutama dalam olahraga basket dan voli, acara ini membuka peluang bagi siswa-siswa berbakat untuk menunjukkan keahlian dan semangat kompetitif mereka di lapangan.
		
		Selain itu, kami juga menyajikan serangkaian kegiatan menarik lainnya yang melibatkan beragam bakat, mulai dari lomba seni, pertunjukan musik, hingga ajang kecerdasan dan pengetahuan. Semua ini dirancang untuk menciptakan pengalaman berkesan bagi setiap peserta dan penonton yang hadir.
		
		Mengusung semangat kebersamaan, kejujuran, dan sportivitas, Event Cup Sekolah Notre Dame bukan hanya sekadar kompetisi, tetapi juga sebuah perayaan keberagaman dan prestasi di kalangan siswa. Kami mengundang Anda untuk bergabung, merayakan prestasi, dan mendukung semangat fair play yang menjadi inti dari acara ini.
		
		Tunjukkan dukungan Anda, rasakan semangat kompetisi yang menyala, dan jadilah bagian dari perayaan luar biasa di Cup Sekolah Notre Dame. Mari bersama-sama menciptakan kenangan tak terlupakan!`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			IsOnline:       null.NewBool(false, true),
			Tags:           []string{"olimpiade", "kompetisi", "olahraga", "seni", "prestisius"},
			FeaturedImages: []string{"https://ytknews.id/wp-content/uploads/2022/03/WhatsApp-Image-2022-03-12-at-10.51.40.jpeg"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("Jiexpo Kemayoran - Jakarta Utara", true),
			Category:        []string{"Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 6, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 6, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 7, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 8, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("40.7128", true),
			Longitude:       null.NewString("-74.0060", true),
			Title:           null.NewString("BIGBANG Jakarta Matsuri", true),
			Description: null.NewString(`Konnichiwa Minasan! Mempersembahkan sebuah rangkaian acara kepada Sahabat Jejepangan 11 Days Japan Festival
			BIGBANG Jakarta Matsuri
			22 Desember 2023 - 1 Januari 2024
			Kolaborasi antara RAF Creative dengan Expo Indo.
			
			Jiexpo Kemayoran - Jakarta
			
			Contents :
			1. Japanese Street food Bazaar
			2. Japanese bazaar 
			3. Cosplay Competition
			4. Coswalk Competition 
			5. Idol performances
			4. DJ Live show
			5. Wibu Berkaraoke
			6. Japan Rock Band
			7. Yosakoi
			8. J - Song Karaoke Performances
			9. Harajuku Fashion Week
			10. Cosplay Artist Alley
			11. Music concert - Indo local band & soloist
			
			Info bazaar & sponsorship
			Hotline whatsapp : 0858-9181-6962
			
			Pendaftaran performer
			Denis Fitz : 0813-8768-1373
			Zee : 0895-1810-7884`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			IsOnline:       null.NewBool(false, true),
			Tags:           []string{"cosplay", "japan", "bazaar", "bigbangjakartamatsuri", "streetfood"},
			FeaturedImages: []string{"https://jakarta-tourism.sgp1.cdn.digitaloceanspaces.com/events/Y6QGKXgLrZIh6ngGZzJYDklWSDJ3X8-metaYmJmLWprdC1mZXN0LTIwMjMtaW1nLWdhbC5wbmc=-.png"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Entertainment & Performance", "Sport"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 8, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 9, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Now(), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("0.0", false),
			Longitude:       null.NewString("0.0", false),
			Title:           null.NewString("Liga Mahasiswa Basketball blibli.com West Java Conference Season 7", true),
			Description:     null.NewString(`Ikuti keseruan Liga Mahasiswa Basketball blibli.com West Java Conference Season 7! Event ini akan berlangsung dari tanggal 8 hingga 15 Juli 2019 di GOR Pajajaran, Bandung. Anda dapat menonton pertandingan secara langsung atau melalui live streaming di YouTube. Dukung tim kampus Anda dan saksikan siapa yang akan menjadi juara musim ini. Jangan lewatkan aksi-aksi seru dari para pemain basket mahasiswa! Catatan: Untuk informasi lebih lanjut dan live streaming, kunjungi www.youtube.com/ligamahasiswa.`, true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"LIMABasketball", "LigaMahasiswa", "AwalMasaDepan", "LIMA", "studentathlete", "Mahasiswa", "basketball", "basket", "college", "indonesia", "bliblidotcom"},
			Location:        null.NewString("GOR Pajajaran, Bandung", true),
			FeaturedImages:  []string{"https://ngobrolbasket.files.wordpress.com/2015/08/basket-lands.jpg"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 8, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 9, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Now(), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("0.0", false),
			Longitude:       null.NewString("0.0", false),
			Title:           null.NewString("Festival Musik Rock Jakarta 2024", true),
			Description: null.NewString(`Festival Musik Rock Jakarta 2024: Nikmati Sensasi Rock Terbaik di Akhir Tahun!

			Apakah Anda penggemar musik rock? Jika ya, maka Anda tidak boleh melewatkan Festival Musik Rock Jakarta 2023, yang akan digelar pada tanggal 27-28 Desember 2023 di Gelora Bung Karno, Jakarta. Festival ini akan menjadi pesta musik rock terbesar di Indonesia, yang akan menampilkan lebih dari 50 band rock papan atas dari dalam dan luar negeri.
			
			Anda akan dapat menyaksikan penampilan spektakuler dari band-band legendaris seperti Slank, Superman Is Dead, Burgerkill, Muse, Linkin Park, Metallica, dan Guns N' Roses, yang akan membawakan lagu-lagu hits mereka di atas panggung. Anda juga akan dapat menemukan band-band baru dan berbakat yang siap mengguncang dunia`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			IsOnline:       null.NewBool(false, true),
			Tags:           []string{"rock, festival, musik, konser"},
			Location:       null.NewString("Jiexpo Kemayoran - Jakarta", true),
			FeaturedImages: []string{"https://awsimages.detik.net.id/community/media/visual/2023/06/02/poster-konser-one-ok-rock-di-jakarta.jpeg?w=1200"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Art & Culture, Sport"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 9, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 9, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("PORSENI (Pekan Olahraga dan Seni)", true),
			Description: null.NewString(`Holla mahasiswa dan mahasiswi NGALAM dimanapun kalian berada !!!

			HMTK ITN Malang mempersembahkan PORSENI 6.0 yang pastinya sayang banget untuk dilewatkan, yuk ceki ceki !!
			
			Pada Porseni fantastic 6.0 kali ini ajang perlombaan yang diadakan adalah :
			
			1. Theater 🎬
			
			Biaya pendaftaran : 200k (Grup)
			
			2. E-Sport (Mobile Legend) 🎮
			
			Biaya pendaftaran : 150k
			
			3. Futsal ⚽
			
			Biaya pendaftaran : 350k
			
			4. Badminton🏸
			
			Biaya pendaftaran :
			
			- Ganda : 250k
			
			- Tunggal : 170k
			
			5. Basket 🏀
			
			Biaya pendaftaran :
			
			- Putra ⛹‍♂ : 450k
			
			- Putri ⛹‍♀ : 370k`, true), // Please add a description if available
			Location:       null.NewString("Institut Teknologi Nasional Malang", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"porseni", "itn", "semarang", "olahraga", "basket", "futsal", "theater", "badminton", "esport"},
			FeaturedImages: []string{"https://eventkampus.com/event/detail/3109/porseni-pekan-olahraga-dan-seni#!"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Education & career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Tech Outlook 2024 - Elevating Recruitment in the Cloud Era", true),
			Description: null.NewString(`Seminar Tech Outlook 2024
		
		Tahun 2024 akan menjadi tahun yang penuh ketidakpastian dalam perekrutan talenta. Tekanan yang tinggi di pasar kerja mendorong pemimpin bisnis untuk mencari solusi inovatif seperti perekrutan dengan kecerdasan buatan (AI) atau menggunakan platform komputasi awan yang efisien.
		
		Alibaba Cloud & Jobstreet by SEEK mengundang kamu bergabung di Seminar Tech Outlook 2024 yang akan diisi narasumber : Muhammad Rohibun ( Solution Architect Alibaba Cloud Indonesia )
		
		Acara : Kamis, 14 Des 2023
		Jam : 12.00 – 16.00 WIB
		Tempat : Bandung
		Link RSVP 👉️ https://zfrmz.com/PVNqoQ7n7WBKTQMq9k4N
		
		Kuota terbatas, segera daftarkan diri kamu di acara ini ( FREE ).`, true),
			Location:       null.NewString("Institut Teknologi Bandung, Kab. Bandung", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"ITB", "tech", "seminar", "AI", "komputasiawan", "cloudcomputing"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4486-tech-outlook-2024-elevating-recruitment-in-the-cloud-era.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Education & career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("JOB FAIR BERSAMA PUSAT KARIR PERGURUAN TINGGI SOLORAYA 2023", true),
			Description: null.NewString(`JOBFAIR AKBAR, TERBESAR, DAN TERLENGKAP DI SOLORAYA
		
		DI USUNG LEBIH DARI 20 KAMPUS TERNAMA
		
		DIIKUTI PULUHAN PERUSAHAAN DAN RIBUAN LOWONGAN PEKERJAAN
		
		
		
		JOB FAIR BERSAMA
		
		PUSAT KARIR PERGURAN TINGGI
		
		SOLORAYA 2023
		
		Acara ini didukung oleh kampus merdeka yang bertujuan untuk membantu mahasiswa dan alumni Universitas dalam mencari pekerjaan atau magang. Terdapat lebih dari 50 perusahaan yang akan berpartisipasi dalam acara ini, termasuk perusahaan-perusahaan ternama seperti Microsoft, Google, dan Amazon. Selain itu, acara ini juga akan diisi dengan seminar dan workshop yang akan membahas berbagai topik menarik seputar dunia kerja, seperti cara membuat CV yang menarik, tips wawancara kerja, dan banyak lagi.`, true),
			Location:       null.NewString("Gedung Grha Soloraya (Eks Gedung Bakorwil) Surakarta, Kota Surakarta", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"jobfair", "edukasi", "kampusmerdeka"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4438-job-fair-bersama-pusat-karir-perguran-tinggi-soloraya-2023.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},

		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Education & career"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 14, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Training of Trainers #3, In-person Course Jakarta", true),
			Description: null.NewString(`Asosiasi Profesional Privasi Data Indonesia, yang didukung oleh Schinder Law Firm, mempersembahkan Training of Trainers, Kursus Tatap Muka #3, Rabu-Kamis, 28-29 Februari, pukul 13:00-17:00 WIB, Jakarta.

			Metodologi Pelatihan
			
			Empat langkah menyoroti pentingnya mempelajari modul, menyiapkan dan menyampaikan presentasi secara aktif, berpartisipasi dalam umpan balik rekan, dan melakukan sesi presentasi menyeluruh.
			
			Studi Komprehensif dan Pengetahuan Terkini:
			Teliti semua 10 modul Perlindungan Data Pribadi DPO Online Course.
			Tetap terkini dengan studi kasus internasional, pengetahuan teoritis, dan praktik terbaik.
			Tekankan pemahaman terhadap GDPR dan pengetahuan menyeluruh tentang persyaratan perlindungan data.
			Persiapan Aktif dan Penyampaian Presentasi:
			Sampaikan presentasi berdasarkan modul yang ditugaskan (ditugaskan pada hari pelatihan).
			Libatkan diri dalam sesi Tanya Jawab dengan profesor dan peserta lain untuk membentuk lingkungan belajar kolaboratif.
			Umpan Balik dan Kolaborasi Rekan:
			Berikan umpan balik konstruktif pada presentasi peserta lain untuk mendukung perkembangan mereka.
			Sesi Presentasi Komprehensif:
			Alihkan 15 menit untuk presentasi modul.
			Selenggarakan sesi tanya jawab selama 15 menit setelah setiap presentasi.
			Kriteria Penilaian
			
			Untuk menjadi pelatih bersertifikasi, Anda akan dinilai berdasarkan 10 kriteria berikut:
			
			Pengetahuan Subjek
			Konten Pelatihan
			Keterampilan Komunikasi
			Keterampilan Fasilitasi
			Adaptabilitas
			Keterlibatan Peserta
			Teknik Pelatihan
			Umpan Balik dan Evaluasi
			Manajemen Waktu
			Bahasa Tubuh
			Nilai kelulusan minimum adalah 85% untuk menjadi pelatih DPO bersertifikat oleh APPDI terbaik.
			
			Khusus untuk APPDI CDPO dan Anggota. Amankan tempat Anda sebelum 28 Januari. Bergabunglah dengan Komunitas Pelatih APPDI.
			
			Jadilah kekuatan perlindungan data dan memberdayakan orang lain!`, true),
			Location:       null.NewString("Centennial Tower 36th Floor, Jl.Jend Gatot Subroto Kav. 24-25, Jakarta Indonesia Jakarta Selatan, Jakarta 12930", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"DPO", "edukasi", "data", "course", "pelatihan", "kursus"},
			FeaturedImages: []string{"https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F644815509%2F572826263471%2F1%2Foriginal.20231120-024306?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C2160%2C1080&s=b05c90f991721ed07c8626f3ff5634e4"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Prof. Dr. IBR Supancana", "Prof. Abu Bakar Munir"},
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("JOB FAIR BERSAMA PUSAT KARIR PERGURUAN TINGGI SOLORAYA 2023", true),
			Description: null.NewString(`JOBFAIR AKBAR, TERBESAR, DAN TERLENGKAP DI SOLORAYA
		
		DI USUNG LEBIH DARI 20 KAMPUS TERNAMA
		
		DIIKUTI PULUHAN PERUSAHAAN DAN RIBUAN LOWONGAN PEKERJAAN
		
		
		
		JOB FAIR BERSAMA
		
		PUSAT KARIR PERGURAN TINGGI
		
		SOLORAYA 2023
		
		Acara ini didukung oleh kampus merdeka yang bertujuan untuk membantu mahasiswa dan alumni Universitas dalam mencari pekerjaan atau magang. Terdapat lebih dari 50 perusahaan yang akan berpartisipasi dalam acara ini, termasuk perusahaan-perusahaan ternama seperti Microsoft, Google, dan Amazon. Selain itu, acara ini juga akan diisi dengan seminar dan workshop yang akan membahas berbagai topik menarik seputar dunia kerja, seperti cara membuat CV yang menarik, tips wawancara kerja, dan banyak lagi.`, true),
			Location:       null.NewString("Gedung Grha Soloraya (Eks Gedung Bakorwil) Surakarta, Kota Surakarta", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"jobfair", "edukasi", "kampusmerdeka"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4438-job-fair-bersama-pusat-karir-perguran-tinggi-soloraya-2023.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 13, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 17, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Charity Children Camp 2023 - SERENITY", true),
			Description: null.NewString(`🪐 FKIK Atma Jaya Presents: Charity Children Camp 2023 - SERENITY 🪐
			Sesuai dengan nama serenity, mahasiswa FKIK Atma Jaya berharap untuk bisa memberikan kebahagiaan, ketenangan, dan kedamaian untuk pesertanya. 
			Charity Children Camp tahun ini memiliki tujuan untuk menyediakan sarana edukasi dan dukungan psikososial yang diharapkan dapat memotivasi dan membangkitkan semangat hidup anak-anak penderita Down Syndrome. Selain itu, dengan adanya acara ini diharapkan dapat meruntuhkan stigma masyarakat mengenai Down Syndrome dan mengenal mereka lebih baik dengan cara berinteraksi serta bermain bersama mereka. Acara ini akan dilaksanakan pada 18-19 November, tertutup untuk mahasiswa FKIK Atma Jaya.
			Akan tetapi, Charity Children Camp memberikan kesempatan untuk masyarakat dalam mendukung anak-anak Down Syndrome dari Rumah Ceria Down Syndrome - POTADS sebagai Non-Governmental Organization yang bekerja sama dengen FKIK Atma Jaya, dengan cara berdonasi. Jika saudara/i ingin berdonasi, link dapat diakses melalui http://bit.ly/3FQd59t atau pada barcode yang tersedia pada poster. 
			“We have the same wants and dreams as everyone else. We can do anything anyone else can do. We are more alike than we are different.” says Manager of Advocacy at the National Down Syndrome Society`, true),
			Location:       null.NewString("Wisma Kinasih, Depok", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"charity", "downsyndrome", "ayoberbagi", "FKIKAtma", "atmajaya", "donasi"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/22/4480-charity-children-camp-2023-serenity.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Dr. dr. Bina Akura Sp.A(K)"},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 13, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 17, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Charity Children Camp 2023 - SERENITY", true),
			Description: null.NewString(`🪐 FKIK Atma Jaya Presents: Charity Children Camp 2023 - SERENITY 🪐
			Sesuai dengan nama serenity, mahasiswa FKIK Atma Jaya berharap untuk bisa memberikan kebahagiaan, ketenangan, dan kedamaian untuk pesertanya. 
			Charity Children Camp tahun ini memiliki tujuan untuk menyediakan sarana edukasi dan dukungan psikososial yang diharapkan dapat memotivasi dan membangkitkan semangat hidup anak-anak penderita Down Syndrome. Selain itu, dengan adanya acara ini diharapkan dapat meruntuhkan stigma masyarakat mengenai Down Syndrome dan mengenal mereka lebih baik dengan cara berinteraksi serta bermain bersama mereka. Acara ini akan dilaksanakan pada 18-19 November, tertutup untuk mahasiswa FKIK Atma Jaya.
			Akan tetapi, Charity Children Camp memberikan kesempatan untuk masyarakat dalam mendukung anak-anak Down Syndrome dari Rumah Ceria Down Syndrome - POTADS sebagai Non-Governmental Organization yang bekerja sama dengen FKIK Atma Jaya, dengan cara berdonasi. Jika saudara/i ingin berdonasi, link dapat diakses melalui http://bit.ly/3FQd59t atau pada barcode yang tersedia pada poster. 
			“We have the same wants and dreams as everyone else. We can do anything anyone else can do. We are more alike than we are different.” says Manager of Advocacy at the National Down Syndrome Society`, true),
			Location:       null.NewString("Wisma Kinasih, Depok", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"charity", "downsyndrome", "ayoberbagi", "FKIKAtma", "atmajaya", "donasi"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/22/4480-charity-children-camp-2023-serenity.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Dr. dr. Bina Akura Sp.A(K)"},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("PROGRAM ASISTENSI SEKOLAH: OPEN DONASI", true),
			Description: null.NewString(`‼PROGRAM ASISTENSI SEKOLAH: OPEN DONASI ‼ Halo, Sobat PAS! 🙌🏻
			Kami dari Program Asistensi Sekolah (PAS) dengan senang hati mengajak teman-teman semua untuk ikut berpartisipasi dalam donasi untuk kegiatan kami yang bertujuan membantu anak-anak di SDN Kebon Kopi Bogor dan SDN Pancoranmas 3 Depok menjadi lebih baik. Kebaikan dalam bentuk apapun akan sangat berarti untuk mereka yang ada di sekolah😊
			DONASI BARANG: 
			📦Dikumpulkan dalam wadah 📍Titik Penyerahan: FIB UI atau Asrama UI
			☎: Yelena (082246266727)
			DONASI UANG: 
			💰Transfer via Bank BNI (1230498153) a.n. Najwa Salsabila
			Mari bersama-sama memberikan harapan dan masa depan yang lebih baik untuk anak-anak pada sekolah dasar. Setiap kontribusi Anda sangat berarti bagi mereka.`, true),
			Location:       null.NewString("FIB UI, Kota Depok", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"charity", "downsyndrome", "ayoberbagi", "FKIKAtma", "atmajaya", "donasi"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4469-program-asistensi-sekolah-open-donasi.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Sport"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 5, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("FUN WALK 90's", true),
			Description: null.NewString(`Fun Walk 90’s
			Local Brand Fashion Market & Kuliner
			Music Performance
			Result
			KESAN, Keseruan yang berbeda di tengah kebosanan gelaran event olahraga santai
			EXPOSURE, Digital Traffic meningkat karena peserta Upload keseruan saat Event Berlangsung
			TRUST, Kepercayaan masyarakat dan sponsor terhadap STIE TOTALWIN meningkat
			Dengan target minimal 5000 peserta akan mengikuti FUN WALK 90's yang digelar oleh STIE TOTALWIN. Adapun ragam hiburan selain keseruan fashion para peserta diantaranya Local Brand Fashion Market 90an dan puluhan kuliner & jajanan sepanjang Jalan Suratmo, Kota Semarang.`, true),
			Location:       null.NewString("Kampus STIE Totalwin, Kota Semarang", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"olahraga", "kuliner", "jalan-jalan", "hiburan"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4454-fun-walk-90s_1.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sport", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 11, 10, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 11, 17, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("UI Battlegrounds 2024", true),
			Description: null.NewString(`PENDAFTARAN UI BATTLEGROUNDS EKSTERNAL DIBUKA!!

			Calling out Battlers in Indonesia!
			
			Hello esports maniacs!
			
			Pertandingan besar yang mempertandingkan universitas-universitas di Indonesia akan segera dimulai. Persiapkan tim dan anggota terbaikmu. Ingin menunjukkan kemampuan universitas kamu kepada semua orang? Hanya tim mu yang paling menguasai game tersebut? Ingin melihat pertandingan yang seru dan menegangkan? Ini saat yang tepat untuk para Battlers menunjukkan kekuatan tim dari universitas kamu dengan semangat sportivitas!
			
			Yuk, daftarkan tim kalian dan bersiaplah untuk bertanding bersama seluruh Battlers dari universitas yang ada di Indonesia. Pendaftaran batch 3 untuk pertandingan eksternal UI Battlegrounds 2021 resmi dibuka mulai pada tanggal 8 hingga 17 November 2021. Daftarkan melalui tautan pada bio linimasa UI Battlegrounds 2021! Pertandingan akan dimulai pada 22 November hingga 12 Desember 2021. Wah dikit lagi, kan?
			
			💰 FREE
			
			Tunggu apa lagi! Langsung daftarkan tim kamu pada:
			https://linktr.ee/uibg.2021
			https://linktr.ee/uibg.2021
			https://linktr.ee/uibg.2021
			
			Jangan lupa untuk pantau linimasa kami agar tidak tertinggal keseruan yang ada di UIBG tahun ini! Kami tunggu tim terbaik dari universitas kalian di arena pesta tahun ini!
			
			
			Linimasa:
			Instagram: @uibg.2021
			Tiktok: @uibg.2021
			LINE: @125nizth
			Youtube: UI BATTLEGROUNDS`, true),
			Location:       null.NewString("FIB UI, Kota Depok", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"esport", "dota2", "mobilelegends", "PUBG", "Valorant", "esportUI"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/_thumbnail/600x600/ui-battlegrounds-2021-feeds-ig.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Anisa Rahim - Onic Esport"},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("IBSC 2024 BUSINESS CASE COMPETITION", true),
			Description: null.NewString(`Salam, para penggemar bisnis!
			Acara Lomba Kasus Bisnis IBSC 2021 merupakan kompetisi tahunan terbesar yang kini dibuka untuk pendaftaran. Kolaborasi ini melibatkan Garena Indonesia, pengembang game terkemuka di Asia Tenggara, dengan tema "Bagaimana Menciptakan Manajemen Pemain yang Kuat". Lomba ini mengundang mahasiswa aktif program diploma dan sarjana untuk menantang diri mereka sendiri dan mendapatkan pengalaman baru.
			
			Panggilan kepada semua mahasiswa aktif program diploma dan sarjana yang ingin menantang diri mereka sendiri dan mendapatkan pengalaman baru melalui IBSC! Acara ini merupakan peluang untuk mengembangkan keterampilan bisnis dan memahami aspek manajemen pemain dalam industri game.
			
			HADIAH BESAR senilai $2400 USD menanti Anda! Apalagi, penawaran fase Early Bird TERBATAS, jadi segera daftarkan tim Anda SEGERA!!
			Daftar melalui:
			https://bit.ly/IBSC2021Registration
			
			Untuk informasi lebih lanjut dan pertanyaan lebih lanjut, silakan hubungi:
			Sherin +62 81261262947
			Rizka +62 82260693977`, true),
			Location:       null.NewString("Persada Indonesia University, Jakarta Pusat", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"businesscase", "garena", "lomba", "kompetisi"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/_thumbnail/600x600/p.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Garena Indonesia"},
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sport", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("DEPARTEMEN OLAHRAGA PEMERINTAHAN MAHASISWA UNIVERSITAS WIDYATAMA MEMPERSEMBAHKAN WEST 2024", true),
			Description: null.NewString(`DEPARTEMEN OLAHRAGA PEMERINTAHAN MAHASISWA UNIVERSITAS WIDYATAMA MEMPERSEMBAHKAN WEST 2021

			Hello Everyone‼️
			.
			Departemen Olahraga Pemerintahan Mahasiswa Universitas Widyatama mengadakan kompetisi EKSTERNAL WEST 2021. Terbuka untuk seluruh mahasiswa di Indonesia. Serta menangkan total hadiah hingga jutaan rupiah!
			Simak kelanjutannya dibawah ini yaa!
			.
			📅 SAVE THE DATE
			Pendaftaran:
			Mobile Legends : 25 Juni 2021 - 1 Agustus 2021
			PUBG Mobile : 25 Juni 2021 - 8 Agustus 2021
			Pelaksanaan:
			Mobile Legends : 3-7 Agustus 2021
			PUBG Mobile : 10-14 Agustus 2021
			💰REGISTRATION FEE
			Rp. 100.000/team
			.
			Contact Person:
			Mobile Legend : 0852 1777 6975 (Dian)
			PUBGM : 0821 1528 8610 (Alfan)
			Line : @984ksret
			
			
			Untuk info lebih lanjut, cek akun sosial instagram kami ya!
			Instagram : utama.esport
			
			
			Akun sosial media lain
			TikTok : utama.esport
			Youtube : Utama E-Sport
			Facebook : Utama E-Sport
			
			Focus On Your Strategy!`, true),
			Location:       null.NewString("Kota Adm. Jakarta Pusat", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"esport", "mobilelegends", "PUBG", "kompetisi"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/pamflet-story.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sport", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 12, 4, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 12, 10, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 12, 20, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("SMANDU Cup 2020 “Robustrophica”", true),
			Description: null.NewString(`Kami mengajak muda-mudi pemberani di seluruh kota untuk ikut berpartisipasi dalam kejuaraan tahunan SMA NEGERI 52 JAKARTA. Tunggu apalagi? Ayo segera daftarkan dirimu!

			Daftar Kompetisi : 
			
			Futsal SMP, SMA
			Mobile Legend
			Kobar
			PES
			Ratoeh Jaroe
			Kir
			Kobar
			Badminton
			Volley
			Basket
			Ball
			Timeline :
			
			Buka Pendaftaran : 12 Desember 2019 – 10 Januari 2020
			Pertemuan teknis : 10 Januari 2020
			Pembukaan : 18 Januari 2020
			Registrasi :
			Senin – Jumat : 15.30 s.d 17.00
			Sabtu : 10.00 s.d 13.00
			Kompetisi : 20 – 21 Februari 2020`, true),
			Location:       null.NewString("SMA NEGERI 52 JAKARTA", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"esport", "olahraga", "Smandu cup", "futsal", "mobile legends", "basket", "badminton", "volley"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/pamflet-story.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("KEMENTERIAN PENDIDIKAN DAN KEBUDAYAAN BEM UMM PROUDLY PRESENT UNIVERSITY OPINION WRITING COMPETITION", true),
			Description: null.NewString(`Hello Agent of Change👋🏻 Hari Pendidikan Nasional akan segera tiba, KEMENDIKBUD BEM UMM mengadakan Lomba Menulis Opini yang bisa banget kalian ikuti loh✨🤩

			✍🏻UNIVERSITY OPINION WRITING COMPETITION✍🏻
			
			📃TEMA 
			"Pendidikan Masa Kini Dalam Situasi Pandemi COVID-19"
			
			📃Persyaratan
			a. Mahasiswa aktif PTN/PTS seluruh Indonesia (Jenjang D3/D4/S1) 
			b. Terdiri dari 1 orang (Individu)
			c. Karya bersifat Orisinal
			
			💵Biaya Pendaftaran : 25K
			
			📃Timeline : 
			📌Pendaftaran dan Pengumpulan Karya : 22 April - 5 Mei 2021
			📌Penjurian : 6 - 9 Mei 2021
			📌Pengumuman : 10 Mei 2021
			
			Link Pendaftaran :
			https://bit.ly/32as0HN
			
			Link Ketentuan Lomba :
			https://bit.ly/SK_lomba_opini
			
			Hadiah
			🏆Juara 1 : Uang Pembinaan + E-Sertifikat
			🏆Juara 2 : Uang Pembinaan + E-Sertifikat
			🏆Juara 3 : Uang Pembinaan + E-Sertifikat
			💳 Peserta : E-Sertifikat
			
			
			Yuk tunggu apalagi? segera daftarkan dirimu untuk mengikuti dan berpartisipasi dalam LOMBA MENULIS OPINI🤩⚡`, true),
			Location:       null.NewString("Universitas YARSI, Jakarta Pusat", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"menulis", "essay", "opini", "kompetisi"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/_thumbnail/600x600/whatsapp-image-2021-04-28-at-114907.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Art & Culture", "Charity", "Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Charity and Cultural Night AMSA-UGM 2022", true),
			Description: null.NewString(`[Charity and Cultural Night AMSA-UGM 2022]
			✨Hello, good people!✨
			
			Apa itu CCN AMSA-UGM? Charity and Cultural Night (CCN) AMSA-UGM merupakan acara yang diselenggarakan oleh AMSA-UGM setiap tahunnya dalam bentuk pentas teater. Tahun ini, CCN AMSA-UGM 2022 akan disajikan dalam bentuk film musikal yang ditayangkan baik secara offline maupun online lohh 🤩
			
			
			
			CCN AMSA-UGM 2022 mempersembahkan film musikal berjudul "Kali Kedua"
			
			•Kesempatan untuk Berubah•
			
			Karya Nayla Eliza H. & Michelle Joviany.
			
			
			
			Seluruh pendapatan dari penjualan tiket CCN AMSA-UGM 2022 akan didonasikan untuk Panti Asuhan Bina Siwi, Yogyakarta 🥰
			
			
			
			Adapun tiket yang tersedia untuk saat ini, yaitu:
			
			📽 KAMA 1 [Online]
			
			🗓 26 Juni 2022
			
			⏰ 18.00 WIB (sesi 1) 
			
			20.30 WIB (sesi 2)`, true),
			Location:       null.NewString("Cinepolis Lippo Plaza Yogyakarta", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"teater", "pentas", "ugm", "charity"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/21/_thumbnail/600x600/4329-charity-and-cultural-night-amsa-ugm-2022.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 18, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("OPENING GEBYAR NUSANTARA 2021", true),
			Description: null.NewString(`Salam Budaya!

			Apa sih Gebyar Nusantara itu?
			Gebyar Nusantara (@gebyarnusantaraipb) merupakan rangkaian acara kebudayaan terbesar di IPB University yang dilaksanakan oleh Kementrian Seni dan Budaya BEM KM IPB
			
			Gebyar Nusantara 2021 mengangkat tema “Abhipraya Palawa Nusantara” yang memiliki arti harapan nusantara yang bersemi kembali. Gebyar Nusantara diikuti oleh 26 Organisasi Mahasiswa Daerah dan akan dimeriahkan dengan berbagai rangkaian acara mulai dari Opening, Bincang Budaya, Pemilihan Putera-Puteri OMDA, Video Creative Competition, dan Malam Puncak.
			Saksikan kemeriahannya dengan special performance Kisah Abhipraya dan mari bersama turut lestarikan warisan budaya nusantara!
			`, true),
			Location:       null.NewString("Kota Bogor", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"nusantara", "budaya", "seni"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/21/_thumbnail/600x600/4329-charity-and-cultural-night-amsa-ugm-2022.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Art & Culture", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 18, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 18, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("GREMOF TALENT SHOW GENERASI MUDA CINTA BANGSA", true),
			Description: null.NewString(`__HIMPRO PKK Proudly Present_
			✨GREMOF TALENT SHOW✨
			"GENERASI MUDA CINTA BANGSA"
			
			Hai Guys..
			Apa kabar Pelajar dan Mahasiswa seluruh Indonesia? Bagaimana sekolah atau kuliah di masa pandemi ini? Masih semangat ya
			Ada kabar gembira untuk kalian semua🤗
			
			﹋﹋﹋﹋﹋﹋﹋﹋﹋﹋﹋﹋
			
			Bagi kalian Pelajar Seluruh Indonesia (SMP, SMA, SMK, Perguruan Tinggi) yang tetap ingin produktif dan memiliki bakat seni atau keterampilan lainnya bisa dengan ikut mendaftarkan dirimu/kelompokmu sekarang !!
			🍂Ayo jangan sampai ketinggalan, dengan mengisi form pendaftaran sesuai kategori yang kalian minati🍂
			🎨 KATEGORI SENI
			http://bit.ly/SeniGTS
			🎬 KATEGORI VIDEO KREASI 
			http://bit.ly/VideoKreasiGTS
			
			Jangan lupa catat TIMELINE⏰
			📝Pendaftaran :
			 1-12 Juni 2021
			📍Technical meeting via zoom :
			 13 Juni 2021
			🎞️ Pengumpulan video :
			 14-24 Juni 2021
			🖊️ Penjurian :
			 25-28 Juni 2021
			🏅 Penayangan video dan pengumuman juara :
			 29-30 Juni 2021
			🏆 Pembagian hadiah :
			 1-3 Juli 2021
			💙 SEMANGAT PRODUKTIF💙
			`, true),
			Location:       null.NewString("UNNES, Semarang", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"nusantara", "budaya", "seni", "makeup", "kompetisi", "talentshow"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/whatsapp-image-2021-06-04-at-32244-pm.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 11, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 11, 18, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Agrifest 4.0", true),
			Description: null.NewString(`Agrifest merupakan festival musik tahunan yang selalu diadakan di Fakultas Pertanian , Universitas Sumatera Utara. Nama Agrifest merupakan festival musik ke-4, setelah sebelumnya ada beberapa nama lain. Agrifest merupakan acara puncak perayaan Ulang Tahun IMASEP Ke-39 dan Agribisnis Ke-59. Sebelumnya ada beberapa kegiatan yang dilaksanakan seperti : pengabdian masyarakat yang diadakan 5 Oktober 2019, kemudian dilanjutkan Agribusiness Competition yang dimeriahkan oleh semua mahasiswa agribisnis untuk tujuan mempererat tali persaudaraan lagi. Kemudian ada kegiatan perayaan Ulang Tahun IMASEP Ke-39 dan Agribisnis Ke-59. Dan acara terakhirnya yaitu AGRIFEST. 
Event Agrifest selalu mengundang artis ibukota dan dimeriahkan oleh band-band lokal medan, seperti di Tahun 2017 kemarin mengundang Fourtwnty, di tahun 2018 mengundang Fiersa Besari x Kerabat Kerja, dan tahun ini mengundang Senar Senja, Dimana lagu Senar Senja yang sedang booming yaitu "Dialog Hujan", "Menua Berdua", "Havanah", "Untuk Yang Baru Saja di Wisuda", "Malu tapi Mau", "Asmara Roda Dua", dan masih banyak lagi. Tidak hanya menampilkan artis saja, juga ada band lokal yang ikut memeriahkan seperti : Hello Benji, The Sugarcane, The Boxquitos, Pesawat Sederhana, dan Not Xmprewell. Ada festival makanan juga yang ikut memeriahkan. Semua Makanan Dijamin Enakkkkk!!!!! 
			`, true),
			Location:       null.NewString("Fakultas Pertanian USU, Kota Medan", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"agrifest", "musik", "konser"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/2/_thumbnail/600x600/agrifest-40.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Senar Senja", "TheBoxQuitos", "The Sugar Cane", "Pesawat Sederhana", "Not Xmprewell"},
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 18, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Malam Puisi Airlangga", true),
			Description: null.NewString(`[HMD SASINDO UNAIR 2019]

			.
			
			DIVISI MINAT DAN BAKAT
			
			mempersembahkan: 
			
			✨Malam Puisi Airlangga✨ 
			
			"Urgensi Puisi Masa Kini” 
			
			.
			
			Mari beramai-ramai kita berdiskusi dan membaca puisi dengan bait bait yang bercerita 
			
			.
			
			Yang diselenggarakan pada:
			
			📅 Sabtu, 21 Desember 2019
			
			🕰 06.00 PM
			
			📍 Hall Lt. 1 Fakultas Ilmu Budaya
			
			.
			
			Penampilan Special : 
			
			Teater Gapus (Ormawa FIB) 
			`, true),
			Location:       null.NewString("Fakultas Ilmu Budaya Universitas Airlangga, Kota Surabaya", true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"agrifest", "musik", "konser"},
			FeaturedImages: []string{"https://eventkampus.com/data/event/poster/3/malam-puisi-airlangga.jpeg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 18, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("NANDAYO japan pop culture week", true),
			Description: null.NewString(`NANDAYO! comeback!! 

			Event yang diselenggarakan oleh Barcode Organizer kali ini akan lebih mengedepankan budaya pop culture Jepang yang sudah pasti jadi favorite kalian semua. 
			
			Di acara NANDAYO! yang kedua ini, mengangkat tema "pop-up yourself", yaitu sebagai wadah expresi para pecinta budaya pop culture Jepang dalam cosplay, fashion, make-up, atau mungkin kecintaan kalian akan sub-culture lainnya dari pop culture Jepang seperti manga, idol group, anisong, music, dan juga aktivitas permainan lainnya. 
			
			Dapatkan update terbaru event "NANDAYO!" di instagram kami : @nandayo.id
			
			Jangan lupa untuk dapatkan tiketmu di Aplikasi GOERS.
			`, true),
			Location: null.NewString(`Senayan Park (SPARK)
			Jl. Gerbang Pemuda No.3, RT.1/RW.3, Gelora, Kecamatan Tanah Abang, Kota Jakarta Pusat, Daerah Khusus, Tanah Abang, Jakarta Pusat Kota, Jakarta, Indonesia
			Pulau Satu`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"jepang", "musik", "konser", "japan", "cosplay", "fashion", "anisong"},
			FeaturedImages: []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702308316-mzPIzZcALZDPlkHzHGHDjNgRQg1nCB5x.jpg?width=1024&quality=90"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 20, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 21, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("BAHTERA VOL. 02", true),
			Description: null.NewString(`Bahtera Vol.2 merupakan event lanjutan dari Bahtera Vol.1 yang dimana HangtuahCup akan kembali mengadakan Pensi Festival Musik sebagai penutupan acara pada tanggal 27 november 2023, event pensi yang diadakan akan mengundang beberapa guest star untuk mengisi event tersebut, pada pensi ini merupakan event yang dijalani oleh siswa siswi SMA HANGTUAH 1 JAKARTA sebagai bentuk memeriahkan hangtuah cup 2024!
			`, true),
			Location: null.NewString(`GOR Seskoal
			RT.1/RW.11, Cipulir, Kby. Lama, Kebayoran Lama, Jakarta Selatan, Jakarta, Indonesia
			GOR SESKOAL`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"konser", "musik", "bahtera", "hangtuah"},
			FeaturedImages: []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1704444134-5FIrt3yyNk7N1QI98jgXw9jklLN0WTxV.PNG?width=1024&quality=90"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Nadin Anizah", "Nadhif Basalamah", "Yahya"},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 20, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 22, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 6, 24, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 6, 24, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Workshop Better Sleep", true),
			Description: null.NewString(`SUDAHKAH TIDUR ANDA BERKUALITAS?

			Tidur yang berkualitas bukan hanya akan mencegah agar anda tidak mengantuk keesokan harinya tetapi juga meningkatkan sistem imunitas, menyeimbangkan hormon, meningkatkan ketajamaen berpikir dan memori serta memulihkan energi dan mencegah penyakit kronis seperti kanker, penyakit jantung, dan diabetes.
			
			Workshop satu hari yang akan mengupas mengenai pola permasalahan tidur anda.
			
			Kenali masalah dan gagngguan tidurmu dan cara yang tepat untuk mengatasinya.
			
			Beberapa hal yang akan dibahas pada workshop ini:
			
			1. Mengetahui pola tidur yang normal dan sehat
			
			2. Mengenal pola tidur yang anda miliki
			
			3. Mengetahui ragam cara natural untuk mengatasi gangguan tidur
			
			4. Mengetahui cara - cara mengatasi gangguan tidur sesuai pola permasalahan tidur anda
			`, true),
			Location:       null.NewString(`Central Park, Jakarta Barat`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"seminar", "workshop", "tidur", "sleep"},
			FeaturedImages: []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1703650368-Eaiba4iQyVkSOiq5lgcWXoTkm6rOLEB1.png?width=1024&quality=90"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 25, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 26, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 7, 2, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("New World Order - The Premier Crypto Talk Experience by Akademi Crypto", true),
			Description: null.NewString(`"New World Order" talk event adalah sebuah acara yang mengajak peserta untuk menjelajahi masa depan dunia crypto secara mendalam. Acara ini akan dilaksanakan pada hari Sabtu, 20 Januari 2024, mulai pukul 18.00 hingga 21.00 di Sutera Hall, Alam Sutera, Tangerang.
			Mari bergabung dalam perjalanan luar biasa ke masa depan crypto pada acara "New World Order" talk event.

Acara ini diselenggarakan pada:

Tanggal dan Waktu: Sabtu, 20 Januari 2024, dari pukul 18.00 hingga 21.00
Tempat: Sutera Hall, Alam Sutera, Tangerang
Persiapkan Diri untuk Perjalanan Kosmik Crypto:
Merupakan malam penuh pembicaraan crypto yang memperluas pikiran, melampaui dunia crypto.

Terhubung dan Merayakan:
Bersiaplah untuk terhubung dan terlibat dalam percakapan dinamis yang merayakan awal dari New World Order of Crypto.
Tempat Terbatas - Pastikan Tempat Anda Sekarang!
Pastikan tempat Anda dengan mendaftar sekarang dan tetap terinformasi dengan informasi terkait acara ini di akun Instagram kami @akademicryptocom.
			`, true),
			Location: null.NewString(`Sutera Hall
			Mall @ Alam Sutera, Jl. Jalur Sutera Bar. No.16, RT.001/RW.004, Panunggangan Tim., Kecamatan Pinang, Tangerang Kota, Banten, Indonesia
			Sutera Hall, Alam Sutera, Tangerang`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"talkshow", "workshop", "crypto"},
			FeaturedImages: []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1703650368-Eaiba4iQyVkSOiq5lgcWXoTkm6rOLEB1.png?width=1024&quality=90"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career", "Art & Culture", "Entertainment & Performance", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 25, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 26, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Nanotenri 2023", true),
			Description: null.NewString(`Festival budaya Jepang yang diselenggarakan oleh ekskul Bahasa Jepang SMA Labschool Cibubur dengan lomba - lomba seperti cosplay, obakeyashiki, hanabi, special performance, karaoke, menggambar manga dan lain-lain. 

			Menampilkan berbagai performance, dari peserta lomba hingga guest star spesial kami! 
			`, true),
			Location: null.NewString(`SD - SMP - SMA LABSCHOOL CIBUBUR
			Jl. Raya Hankam Kampus Labschool No.15-20, Jatiranggon, Kecamatan Jatisampurna, Bekasi Kota, Jawa Barat, Indonesia`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"jepang", "cosplay", "labschool"},
			FeaturedImages: []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1697986623-dgOWDEbq7mkzMlNH3imPFRfx8xcCO8SI.png?width=1024&quality=90"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Elaine"},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity", "Sport"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 25, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 26, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Ayo Donor Darah - Renotop", true),
			Description: null.NewString(`Hi Renotopers!

			Siapa nih yang suka donor darah atau ikut kegiatan senam sehat??
			
			Buruan cek @renotop.id karena akan banyak event-event serta hadiahnya.
			
			Ayo ikutan event-event di @renotop.id & catat tanggalnya ya 🤩
			
			-DONOR DARAH di @renotop.meruya
			📍 Lokasi : Jl. Meruya Selatan, No. 19, Kembangan - Jakarta Barat
			🗓️ Hari/Tanggal : Sabtu, 11 Maret 2023
			🕓 Pukul : 10.00 - 14.00 WIB
			🎁 FREE GIFT 25 Pendonor Pertama
			
			-SENAM SEHAT CERIA di @renotop.raden.patah
			📍 Lokasi : Jl. Raden Patah No.20, RT.001/RW.004, Sudimara Bar., Kec. Ciledug, Kota Tangerang, Banten 15151
			🗓️ Hari/Tanggal : Sabtu, 18 Maret 2023
			🕓 Pukul : 07.00 - 08.00 WIB
			.
			.
			Info Selengkapnya cek ke @renotop.id atau dapat menghubungi:
			📞 Syahrul : 0856-9228-5834⁣⁣
			.⁣⁣
			.
			Renovasi Top? Ya RENOTOP⁣⁣
			✔TOP Pilihan Motifnya
			✔TOP Promonya
			✔TOP Pengirimannya
			`, true),
			Location:       null.NewString(`Jl. Meruya Selatan, No. 19, Kembangan - Jakarta Barat`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"donor", "donor darah", "senam sehat", "proyekrumah", "renovasirumahrenotop"},
			FeaturedImages: []string{"https://renotop.id/img/blogs/donor-darah-di-renotop-meruya-17-juni-2023-Slsfq.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 28, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Carnival Fest Jungleland", true),
			Description: null.NewString(`Halo Sahabat Jungleland

			Nikmati liburan akhir tahun di Carnival Fest Jungleland Sentul yuk...
			
			Bakal banyak keseruan, Music Festival, Cosplay Competition, Anikaraoke, hingga Clothing & Culinary Festival
			
			Ada Guest Star keren dari Kangen Band, JKT48, hingga Erie Suzan-Beniqno dan sederet artis seru lainnya
			
			Nikmati sensasi bermain di 33 wahana seru sekaligus menyaksikan penampilan spesial dari JKT 48 dan Kangen Band.
			
			Jadi buat kalian gak perlu bingung lagi nih buat mengisi liburan akhir tahun
			
			Untuk informasi Bazaar hubungi:
			- Irex 087774794666
			- Zulkifli Chimot 081350117416
			- Alfi 081585548836
			
			Yuk beli tiketnya sekarang di www.ticket.jungleland.id
			
			Buruan sebelum kehabisan
			
			Sampai jumpa di Carnival Fest Jungleland Sentul 26 Desember 2022 sampai 1 Januari 2023
			`, true),
			Location:       null.NewString(`Jungleland Sentul`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"cosplay", "musik", "konser", "JKT48", "akhirtahun"},
			FeaturedImages: []string{"https://images.genpi.co/uploads/arsip/normal/2022/11/30/ilustrasi-menikmati-momen-libur-akhir-tahun-di-c-tfur.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"JKT48"},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Indonesia Makeup Expo (IMAE) Back To Beauty", true),
			Description: null.NewString(`Indonesia Makeup Expo (IMAE) Back To Beauty
			Tanggal : 31 Agustus - 4 September 2022
			Tempat : Atrium Utama, Aeon Mall BSD CITY
			Waktu : 10.00 - 22.00 WIB
			
			Deskripsi Event :
			Calling all Makeup and Beauty Enthusiasts!
			
			Indonesia Makeup Expo (IMAE) Back To Beauty akan hadir kembali di Aeon Mall BSD City. Akan ada lebih dari 35 Brand Makeup, Skincare dan masih banyak lagi.
			
			Kamu bisa berbelanja Makeup dan Skincare favorit kamu dengan Diskon up to 50% dan mendapatkan berbagai FREE Items.
			
			Selain berbelanja, kamu bisa ikutan FREE seminar makeup demo dari Top Professional Makeup Artists seperti Slam Wiyono, Fauzia Hanum, Natcha, Yoga Septa, Karen Shenna. Tidak ketinggalan kamu bisa bertemu dengan Allyssa Hawadi dari Button Scarves Beauty dan Uly Novita, Content Creator yang viral karena bisa menirukan ratusan wajah artis terkenal.
			`, true),
			Location:       null.NewString(`Aeon Mall BSD City`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"makeup", "expo", "aeon", "makeupexpo", "BSD"},
			FeaturedImages: []string{"https://images.genpi.co/uploads/arsip/normal/2022/11/30/ilustrasi-menikmati-momen-libur-akhir-tahun-di-c-tfur.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"JKT48"},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sport", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Nasional Freethrow Basketball Open Championship", true),
			Description: null.NewString(`Nasional Freethrow Basketball Open Championship
			Pengumuman Pemenang : 1 Desember 2020
			
			Timeline : 
			
			Pendaftaran : 5 Oktober -15 November 2020
			Pengiriman Video : 17 - 22 November 2020
			Penilaian : 22 - 29 November 2020
			Pengumuman Pemenang : 01 Desember 2020
			Syarat dan Ketentuan : 
			
			Peserta mendaftar secara online pada link http://bit.ly/FreeThrowChampionship2020 (Link on bio) dan wajib follow official account event @freethrow_championship2020
			Konfirmasi pendaftaran pada CP tertera.
			Membayar biaya registrasi ke alamat rekening yang diberikan oleh CP.
			Melakukan pengambilan video freethrow/ memasukan bola basket sebanyak mungkin ke ring dengan durasi dan ketentuan yang sudah diberikan.
			Mengunggah video pada drive (pastikan bisa diakses oleh publik) dan kirimkan Link drive video pada email, dengan subjek yg sudah ditentukan : freethrowopen2020@gmail.com
			Penilaian juri dilakukan dengan berpedoman pada kriteria penilaian.
			Pemenang lomba akan mendapatkan medali, uang pembinaan, dan piagam perhargaan.
			Pendaftaran :
			
			Biaya Pendaftaran :
			Dewasan : Rp 50.000
			Pelajar : Rp 30.000
			`, true),
			Location:       null.NewString(`Lapangan Budi Indah, Komplek Budi Indah Daan Mogot Km 18`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"basket", "kompetisi", "freethrow"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2020/10/MP-EJ-Permata-Freethrow-Championship-2020-Mahasiswa-Unsoed-UNS-dan-UPI-Copy.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo", "Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 28, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Festival Glodok Kuliner Kota Tua", true),
			Description: null.NewString(`Festival Glodok Kuliner Kota Tua
			Tanggal : 19 - 24 Desember 2023
			Tempat : Halaman Depan Glodok Plaza
			Waktu : 10.00 - 18.00
			
			Deskripsi Event :
			CALLING ALL FOOD TENANT !!!
			
			Glodok Plaza & @jmlorganizer85
			
			Mempersembahkan Festival Glodok Kuliner Kota Tua tempat nya di Out door Glodok Plaza
			
			Datang aja dan langsung pilih kuliner kesukaan kamu di festival Glodok Kuliner Kota Tua,,selain itu kamu bisa nikmati live music, talent perform dan berbagai lomba hingga Cosplay
			
			Jangan lupa yaa Festival Glodok Kuliner Kota Tua 19  24 Desember 2023 DARI JAM 10.00  18.00
			`, true),
			Location:       null.NewString(`Kota Tua, Glodok Plaza`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"festival", "kotatua", "kuliner", "cosplay", "livemusic", "musik"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2023/12/PP-EJ-Festival-Glodok-Kuliner-Kota-Tua-JML-Organizer-85-Copy.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 30, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 30, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Kejar Mimpi UPSKILL V.10", true),
			Description: null.NewString(`Kejar Mimpi UPSKILL V.10
			Tanggal : 9, 14 Desember 2023
			
			Deskripsi Event :
			[KEJAR MIMPI TANGSEL PROUDLY PRESENTS : UpSkill 1.0]
			
			Hello Dream Warriors!!👋🏼
			
			🚀 Yuk, gaspol naik level skill desain kamu! Dateng ke workshop kita yang keren abis!
			
			Kamu bakal dapet ilmu ngebut tentang Design Thinking di hari pertama. Lalu, di hari kedua, kita eksplorasi dunia UI/UX Design yang bikin skill kamu makin keren!
			
			Pastinya, ini bisa jadi kesempatan emas buat kamu para Dream Warriors untuk menambah wawasan, kenalan dengan expert di bidangnya, dan yang pasti kesempatan untuk upgrade kemampuan desain kamu!
			
			Catat tanggalnya, dan jangan sampai kelewatan, ya! 
			
			DESIGN THINKING 101
			Tanggal : 9 Desember 2023
			Tempat : Digital Lounge CIMB Niaga, Central Park
			UI/UX DESIGN 101
			Tanggal : 14 Desember 2023
			Tempat : Digital Lounge CIMB Niaga, Atma Jaya
			`, true),
			Location:       null.NewString(`Digital Lounge CIMB Niaga, Central Park`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"UI", "UX", "Design Thinking", "Seminar", "Workshop", "UI/UX"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2023/12/MP-EJ-UpSkill-1.0-KEJAR-MIMPI-TANGSEL-1-Copy.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{"Aswin Widyastama", "Miqdad Darmawan"},
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Mari Bantu Moestopo Jelajah Nusantara 2024", true),
			Description: null.NewString(`Mari Bantu Moestopo Jelajah Nusantara 2024

			Deskripsi Event :
			❗#BantuMJN2024 ❗
			
			Menurut data dari Riset Kesehatan Dasar (Riskesdas) tahun 2018, prevalensi nasional masalah gigi dan mulut di Indonesia adalah 57,6 persen dan hanya 10,2 persen yang telah mendapatkan pelayanan dari tenaga medis. Menurut laporan Riskesdas Provinsi Maluku Tahun 2018, Provinsi Maluku termasuk ke dalam daerah 3 tertinggi dengan angka prevalensi gigi rusak/berlubang/sakit yaitu sebesar 56,28%.
			
			Yuk bantu kami untuk terwujudnya Senyum Sehat, Senyum Ceria Masyarakat Maluku melalui link berikut: https://kitabisa.com/campaign/yukbantumjn24
			
			Bantuan dari teman-teman sangat berarti untuk terwujudnya Senyum Sehat, Senyum Ceria Masyarakat Maluku #MengabdiDenganHatiBergerakDenganAksi #BantuMJN2024.
			`, true),
			Location:       null.NewString(``, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"charity", "kesehatan", "mulut", "gigi"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2023/12/MP-EJ-Moestopo-Jelajah-Nusantara-2024-Senat-Mahasiswa-FKG-Senat-Mahasiswa.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Mari Bantu Moestopo Jelajah Nusantara 2024", true),
			Description: null.NewString(`Mari Bantu Moestopo Jelajah Nusantara 2024

			Deskripsi Event :
			❗#BantuMJN2024 ❗
			
			Menurut data dari Riset Kesehatan Dasar (Riskesdas) tahun 2018, prevalensi nasional masalah gigi dan mulut di Indonesia adalah 57,6 persen dan hanya 10,2 persen yang telah mendapatkan pelayanan dari tenaga medis. Menurut laporan Riskesdas Provinsi Maluku Tahun 2018, Provinsi Maluku termasuk ke dalam daerah 3 tertinggi dengan angka prevalensi gigi rusak/berlubang/sakit yaitu sebesar 56,28%.
			
			Yuk bantu kami untuk terwujudnya Senyum Sehat, Senyum Ceria Masyarakat Maluku melalui link berikut: https://kitabisa.com/campaign/yukbantumjn24
			
			Bantuan dari teman-teman sangat berarti untuk terwujudnya Senyum Sehat, Senyum Ceria Masyarakat Maluku #MengabdiDenganHatiBergerakDenganAksi #BantuMJN2024.
			`, true),
			Location:       null.NewString(``, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"charity", "kesehatan", "mulut", "gigi", "Maluku"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2023/12/MP-EJ-Moestopo-Jelajah-Nusantara-2024-Senat-Mahasiswa-FKG-Senat-Mahasiswa.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Mari Bantu Moestopo Jelajah Nusantara 2024", true),
			Description: null.NewString(`Mari Bantu Moestopo Jelajah Nusantara 2024

			Deskripsi Event :
			❗#BantuMJN2024 ❗
			
			Menurut data dari Riset Kesehatan Dasar (Riskesdas) tahun 2018, prevalensi nasional masalah gigi dan mulut di Indonesia adalah 57,6 persen dan hanya 10,2 persen yang telah mendapatkan pelayanan dari tenaga medis. Menurut laporan Riskesdas Provinsi Maluku Tahun 2018, Provinsi Maluku termasuk ke dalam daerah 3 tertinggi dengan angka prevalensi gigi rusak/berlubang/sakit yaitu sebesar 56,28%.
			
			Yuk bantu kami untuk terwujudnya Senyum Sehat, Senyum Ceria Masyarakat Maluku melalui link berikut: https://kitabisa.com/campaign/yukbantumjn24
			
			Bantuan dari teman-teman sangat berarti untuk terwujudnya Senyum Sehat, Senyum Ceria Masyarakat Maluku #MengabdiDenganHatiBergerakDenganAksi #BantuMJN2024.
			`, true),
			Location:       null.NewString(``, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"charity", "kesehatan", "mulut", "gigi", "Maluku"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2023/12/MP-EJ-Moestopo-Jelajah-Nusantara-2024-Senat-Mahasiswa-FKG-Senat-Mahasiswa.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sport", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 11, 13, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 11, 15, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Lomba Futsal  STUDENT CUP 2023", true),
			Description: null.NewString(`Lomba Futsal STUDENT CUP 2023
			Kick Off : 6-9 Maret 2023
			Tempat : LAPANGAN FUTSAL UTAMA LT 5
			
			Deskripsi Event :
			Hi GUYS
			
			UNIT KEGIATAN MAHASISWA FUTSAL TAMA JAGAKARSA
			
			_Proudly present_
			
			STUDENT CUP
			Kompetisi futsal untuk SMA/SMK Se-jabodetabek dengan berbagai hadiah menarik
			
			Dengan membawa tema : Futsal for unity
			
			Pendaftaran :
			
			Tanggal Pendaftaran : 26 Desember 2022 – 3 Maret 2023
			Biaya Pendaftaran : Rp.250.000 + 50.000 (uang deposit)
			Technical Meeting : 4 Maret 2023
			Tempat : KAMPUS UNIVERSITAS TAMA JAGAKARSA
			Kuota : 32 Tim
			Sistem Knock Out
			Persyaratan :
			
			Pas foto berwarna 4x6 (2lembar)
			Foto copy kartu pelajar
			Surat pengantantar dari sekolah
			More Information :
			
			YOEL BANGUN : +62 896-9951-1405
			ANDRI : +62 895-1752-9773
			Instagram : utamafutsal_new
			`, true),
			Location:       null.NewString(`Universitas TAMA JAGAKARSA, JALAN TB Simatupang No. 152 Jagakarsa, DKI Jakarta`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"sport", "olahraga", "futsal", "lomba"},
			FeaturedImages: []string{"https://eventsurabaya.net/wp-content/uploads/2022/06/MP-ES-SNBC-2022-MPK-OSIS-SMAN-1-GRESIK-Copy.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sport", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 11, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("SMANSA Futsal And Basketball Competition (SNBC) 2K19", true),
			Description: null.NewString(`SMANSA Futsal And Basketball Competition (SNBC) 2K19

			Tanggal: 2 Agustus – 1 September 2019 (Pelaksanaan)
			Tempat: Wahana Ekspresi Pusponegoro
			
			SMANSA Futsal And Basketball Competition (SNBC) 2K19 adalah sebuah kompetisi olahraga futsal dan basket yang diselenggarakan oleh PK – OSIS SMA NEGERI 1 GRESIK. Event ini berlangsung mulai tanggal 2 Agustus hingga 1 September 2019 di Wahana Ekspresi Pusponegoro. Peserta dapat mendaftar mulai tanggal 1 hingga 24 Juli 2019 dengan biaya pendaftaran yang berbeda untuk kategori futsal dan basket SMA serta SMP. Sebuah pertemuan teknis dijadwalkan pada 28 Juli 2019 untuk memberikan informasi lebih lanjut kepada peserta.

Total hadiah yang diperebutkan mencapai 35 juta rupiah, disertai dengan piala dan sertifikat bagi pemenang. Pendaftaran memiliki batas kuota, sehingga dianjurkan untuk mendaftar secepat mungkin. Event ini menawarkan pengalaman kompetitif yang menarik bagi para pemain futsal dan basket dari berbagai sekolah.
			
			Hadiah:
			
			Total uang tunai 35 juta rupiah
			Piala
			Sertifikat
			Timeline:
			
			Pendaftaran: 1 - 24 Juli 2019
			Technical Meeting: 28 Juli 2019
			Pelaksanaan: 2 Agustus - 1 September 2019
			Pendaftaran:
			
			Biaya Pendaftaran:
			Futsal SMA: Rp 335.000
			Futsal SMP: Rp 315.000
			Basket SMA: Rp 335.000
			Basket SMP: Rp 315.000
			Formulir: Rp 35.000
			Kuota Terbatas
			`, true),
			Location:       null.NewString(`Wahana Ekspresi Pusponegoro`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"sport", "olahraga", "futsal", "lomba", "Basket", "Kompetisi"},
			FeaturedImages: []string{"https://eventsurabaya.net/wp-content/uploads/2022/06/MP-ES-SNBC-2022-MPK-OSIS-SMAN-1-GRESIK-Copy.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sport", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 12, 13, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("INFINITUM CUP Badminton Competition", true),
			Description: null.NewString(`INFINITUM CUP Badminton Competition
			Tanggal : 17 Desember 2022
			Tempat : GOR Cempaka Putih, Jakarta Pusat
			
			Deskripsi Event :
			Hai
			
			Kabar baik nih untuk temen-temen yang masih di jenjang SMA. Kami akan melakukan event Badminton nihh, pada tanggal 17 Desember 202 di Gor Cempaka Putih. Pastinya untuk kalian yang mau ikutan boleh langsung regis melalui link diatas yaa.
			
			Oh iya jangan lupa menangin hadiah nya, ya gais. See you!!!!
			`, true),
			Location:       null.NewString(`GOR Cempaka Putih, Jakarta Pusat`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"sport", "olahraga", "badminton", "lomba", "infinitum", "Kompetisi"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2022/11/MP-EJ-Lomba-Badminton-Infinitum-Copy.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Art & Culture", "Entertainment & Performance", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 11, 20, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 11, 21, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Valskrie Korean Journey with Ilkom UBakrie", true),
			Description: null.NewString(`Prodi S1 Ilmu Komunikasi UBakrie mempersembahkan Event Bergengsi “Valskrie”, dan kali ini Budaya yang ditampilkan adalah Budaya Korea Selatan dengan tema “Korean Journey with Ilkom UBakrie”.

			Akan banyak penampilan menarik, unik, seru yang akan menggemparkan panggung Ilkom Universitas Bakrie.
			
			Ikuti lomba Dance Cover dan Song Cover sesuai petunjuk pada slide kedua. Menangkan hadiah jutaan rupiah dan para Pemenang akan tampil di puncak acara Valskrie pada 23 Desember 2023 di Universitas Bakrie.
			`, true),
			Location:       null.NewString(`Universitas Bakrie, Jakarta`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"sport", "dance", "song", "cover", "korean", "korea", "kpop"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2023/11/MP-EJ-Valskrie-Korean-Journey-with-Ilkom-UBakrie.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo", "Sport", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 21, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("STATE, Student Activities Unit Explore", true),
			Description: null.NewString(`Akhirnya yang ditunggu datang juga. Apa ya? 🤔 Yang ditunggu yaa STATE MAXIMA 2023 dong! 🤩🔥 Maximers siap untuk menjelajahi yang mana nih? 🏃💨

			Have fun and enjoy, Maximers! 🔥
			
			Timeline :
			
			Senin, 18 September 2023
			Ultima Sonora
			Lions Volley
			Ultimags
			J-Café Cosplay
			J-Café Illustration & Visual Novel
			J-Café Music
			J-Café Culture & TCG
			Selasa, 19 September 2023
			Lion Basket
			UMN Documentation
			UMN Taekwondo
			Street Dance
			Capoeira
			Radar UMN
			UESC Scrabble
			UESC Spelling Bee
			UESC Speech
			EUSC Debate
			Rabu, 20 September 2023
			Lions Tennis Meja
			UMN PC Obscura
			UMN TV
			USO
			Ultima Akido
			Trace Reguler
			Trace Ratoh Jatoe
			Kamis, 21 September 2023
			Fortius
			Teater Katak Aktor
			Teater KataK Properti
			Teater Kata Makeup & Costume
			Teater KataK Music
			Lions Futsal
			U-Toys
			Himars
			Jumat, 22 September 2023
			Ikatan Bikers UMN
			Spectre
			Skystar Venture
			KSPM
			Mufomic
			Nusakara
			Senin, 25 September 2023
			Rencang
			Lions Badminton
			Game Development Club
			UMN Robotic
			UMN Juice
			MAPALA
			U-Bix
			POPSICLE UMN
			Selasa, 26 September 2023
			Qorie K-Code Boy Group
			Qorie K-Code Girl Group
			Qorie K-Voice
			Qorie Hantalk
			Rumpin
			ACES
			`, true),
			Location:       null.NewString(`UMN, Gading Serpong`, true),
			AverageRating:  null.NewFloat64(0, true),
			IsFinished:     null.NewBool(false, true),
			Tags:           []string{"sport", "olahraga", "budaya", "esport", "maxima", "UMN", "Gading Serpong"},
			FeaturedImages: []string{"https://eventjakarta.com/wp-content/uploads/2023/09/MP-EJ-MAXIMA-2023-UMN-1-Copy.jpg"},
			Latitude:       null.NewString("0.0", false),
			Longitude:      null.NewString("0.0", false),
			GuestStar:      []string{""},
		},
	}

	return EventsToInsert
}
