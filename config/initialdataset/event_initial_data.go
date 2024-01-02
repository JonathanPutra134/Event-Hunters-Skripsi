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
			PreregisterDate: null.NewTime(time.Date(2024, 12, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 12, 10, 0, 0, 0, 0, time.UTC), true),
			StartEvent:      null.NewTime(time.Date(2024, 12, 20, 0, 0, 0, 0, time.UTC), true),
			EndEvent:        null.NewTime(time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC), true),
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
	}

	return EventsToInsert
}
