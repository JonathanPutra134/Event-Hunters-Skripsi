package dataset

import (
	"time"

	"event-hunters/models"

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
			StarteventDate:  null.NewTime(time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 9, 11, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 1, 30, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.22359", true),
			Longitude:       null.NewString("106.82336", true),
			Title:           null.NewString("Konser ELFAS SINGERS Gempita Cinta", true),
			Description:     null.NewString("Menandai 45 tahun eksistensi di dunia musik Tanah Air, ELFA'S SINGERS bakal menggelar Konser Musik yang diberi tajuk : GEMPITA CINTA. Panggung musik istimewa ini akan digelar Rabu, 20 Desember 2023 di Theater Ciputra Artpreneur, Kuningan, Jakarta Selatan, pada pukul 19.30 - 22.00 WIB. Penonton diajak mengarungi waktu dari masa ke masa bersama ELFA'S SINGERS, tak pelak suguhan konser kolaborasi apik dengan Dian HP ini akan menjadi panggung musik yang istimewa. ELFA'S SINGERS yang beranggotakan: Yana Julio, Lita Zen, Agus Wisman dan Ucie Nurul, bukan sekedar mengajak penonton bernostalgia dengan lagu-lagu hits mereka di masa-masa silam. menyanyi dengan harmonisasi vokal indah di atas panggung, melekat pada ELFA'S SINGERS. Hal tersebut menjadi inspirasi keren untuk penyanyi generasi milenial atau Gen Z. Yuk nonton sama keluarga / pacar / calon pendamping hidup / teman yang kece, ajak semuanya yah, kalian akan merasakan cinta & kebahagiaan. Sampai ketemu di Artpreneur", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"music", "concert", "eventjakarta", "musik"},
			FeaturedImages:  []string{"https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702389367-yAFLAv0mteaSAI9QUFrKfaTUV1Dlu7hr.jpg?width=1024&quality=90", "https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702389402-T7T2674YtR11LJGQVIiKhFg2RpMhTt8m.jpg?width=1024&quality=90", "https://www.goersapp.com/events/konser-elfas-singers-gempita-cinta--konserelfassingers"},
			Image:           null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702389367-yAFLAv0mteaSAI9QUFrKfaTUV1Dlu7hr.jpg?width=1024&quality=90", true),
			GuestStar:       []string{"Dian HP", "The SECIORIAS", "MYLENE", "ELFA'S JAZZ & POP SINGERS"},
		},
		{
			EventcreatorID:  null.NewInt(1, true),
			Location:        null.NewString("Universitas Nasional,  Jl. Sawo Manila, Pasar Minggu, Jakarta Selatan, Jakarta, Indonesia Lantai 1", true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 14, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.28079", true),
			Longitude:       null.NewString("106.83942", true),
			Title:           null.NewString("STAND UP COMEDY LIMA EMPAT ABDEL ACHRIAN", true),
			Description:     null.NewString("\"LIMA EMPAT adalah stand up comedy show pertama cing abdel ( Abdel Achrian ). sebuah cerita perjalanan seorang anak minang yang lahir di pisangan. LIMA EMPAT akan menjadi sebuah pertunjukan perjalanan seorang seniman komedi bernama Abdel Achrian, dari belia hingga sekarang menjelang usia senja. lima empat sebuah angka yang tidak mudah ditempuh begitu saja, melainkan penuh kisah perjuangan, keluh kesah & cucur keringat. sampai jumpa di LIMA EMPAT sebuah angka setelah lima tiga & sebelumlimalima \"", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"standupcomedy", "abdelachrian", "eventjakarta"},
			Image:           null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702779627-2YCwuzytQ3RsT1GerdZ0BfCaABQVk1O7.jpeg?width=1024&quality=90", true),
			GuestStar:       []string{"Abdel Achrian"},
		},
		{
			EventcreatorID:  null.NewInt(1, true),
			Location:        null.NewString("Jakarta, Indonesia,  Setia Budi, Jakarta Selatan Kota, Jakarta, Indonesia", true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 4, 2, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.21363", true),
			Longitude:       null.NewString("106.82645", true),
			Title:           null.NewString("Growth Hacking Summit 2024", true),
			Description:     null.NewString("Bergabunglah bersama kami di acara premier Indonesia bersama lebih dari 500 profesional pemasaran pertumbuhan, pendiri, dan para ahli produk yang memiliki pemikiran sejenis. Growth Hacking Summit 2024 bukan hanya konferensi; ini adalah perpaduan unik antara inovasi, inspirasi, dan wawasan yang dapat diimplementasikan!", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"edukasi", "seanellis", "innovasi"},
			Image:           null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702877844-R1gaXdUdFUAo7MEv3dZcci5Ds8bAm2EX.png?width=1024&quality=90", true),
			GuestStar:       []string{"Sean Ellis"},
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("AYANA Midplaza Jakarta, Kav 10-11 Jalan Jenderal Sudirman Kecamatan Tanah Abang, Daerah Khusus Ibukota Jakarta 10220", true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 5, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 5, 10, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 5, 10, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 5, 10, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 16, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.20901", true),
			Longitude:       null.NewString("106.81971", true),
			Title:           null.NewString("Begin Edu Fair Jakarta (Indonesia)", true),
			Description:     null.NewString("Begin Edu Fair - Jakarta adalah platform bagi pelajar dan orang tua mereka, serta mahasiswa universitas yang sedang mencari program studi internasional, termasuk kursus persiapan, sarjana, magister, MBA, serta program bahasa dan musim panas. Jangan lewatkan kesempatan untuk bertemu dengan calon mahasiswa dan orang tua mereka yang tertarik dengan program di institusi Anda. Anda akan menerima daftar kontak dari mereka yang menunjukkan minat pada institusi Anda.", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"edukasi", "edufair", "educationfair", "internationalstudyprogram"},
			Image:           null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F543006069%2F143238831910%2F1%2Foriginal.20230626-143225?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C1536%2C768&s=1d7bd63cadf03dea1bf1de4d8625494f", true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("Online", true),
			Category:        []string{"Art & Culture", "Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 3, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 17, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("0", true),
			Longitude:       null.NewString("0", true),
			Title:           null.NewString("Together on the Wheel - Pottery Class by Classpop!", true),
			Description:     null.NewString("Kelas Online kerajinan tanah liat yang memuaskan dan menenangkan dipandu oleh seorang instruktur kerajinan tanah liat berpengalaman melalui Classpop! Pelajari berbagai pendekatan dalam kerajinan ini dan ciptakan karya seni unik. Bergabunglah dengan Classpop! untuk \"Bersama di Roda\" dengan Instruktur Michelle dan telusuri teknik-teknik menarik serta gaya seni dalam kelas kerajinan tanah liat yang dirancang untuk pemula maupun pengrajin berpengalaman. Anda akan menciptakan karya kustom yang akan melayani Anda selama bertahun-tahun, semua dengan bantuan instruktur yang ahli dan bersemangat untuk berbagi pengetahuan mereka tentang kerajinan ini dengan siswa seperti Anda. Classpop! menawarkan kelas-kelas kreatif yang penuh kesenangan, seperti tari, lukisan, kerajinan tanah liat, dan lainnya, di lingkungan yang santai dan dipimpin oleh instruktur kelas dunia. Pelajari, ciptakan, dan dapatkan inspirasi dengan memesan kelas hari ini! Untuk informasi tambahan dan tanggal kunjungi Together on the Wheel!", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(true, true),
			Tags:            []string{"pottery", "kerajinan", "tanahliat", "art", "seni"},
			Image:           null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F626529949%2F316402407909%2F1%2Foriginal.20231023-152114?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C2160%2C1080&s=40a43805d992d206c7c72c5d24f0614b", true),
		},
		{
			EventcreatorID:  null.NewInt(1, true),
			Location:        null.NewString("Gambir Expo Kemayoran Jl. Benyamin Suaeb, Pademangan Tim., Kec. Pademangan, Jakarta Utara", true),
			Category:        []string{"Entertainment & Performance", "Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 3, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.14346", true),
			Longitude:       null.NewString("106.84702", true),
			Title:           null.NewString("MOTION IME FESTIVAL", true),
			Description:     null.NewString("Kamu tau gak, kalau banyak kabupaten di Papua masuk dalam daftar daerah tertingal? Hal ini tidak lain karena akses pendidikan yang tidak merata dan Pemberdayaan SDM yang belum optimal! Melalui MOTION IME FEST 2023, berkolaborasi dengan Yayasan Cakra Abhipraya Responsif, memiliki inisiatif untuk membuat sebuah Konser amal, yang seluruh dari hasil penjualan tiketnya nanti akan di donasikan untuk membuat education center, di daerah papua pegunungan, yang di konsep berupa sekolah alam, untuk mendukung peningkatan mutu pendidikan dan pengembangan generasi penerus papua. Program ini akan di jalankan awal tahun 2024 di Distrik walaik, Kabupaten Jayawijaya Provinsi Papua Pegunungan. Selain program pendidikan kami juga akan menjalankan program gizi untuk menurunkan angka stunting. Ayo sebarkan kabar baik ini dan jadi lah bagian dari kebaikan dengan, datang ke acara MOTION IME FEST 2023 2-3 Desember 2023 di Gambir Expo Kemayoran dan nikmati berbagai macam pertunjukan termasuk cosplay, konser dan talkshow", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"cosplay", "hakken", "windahbasudara", "eventcosplay", "charity"},
			Image:           null.NewString("https://ik.imagekit.io/playstagingid/event/banner-1699796688018_0cTNNQadT", true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("Hotel Grand Tjokro, Jalan Daan Mogot No.63 Tj. Duren Utara Kec. Grogol petamburan, Daerah Khusus Ibukota Jakarta 11470", true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 3, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 3, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 5, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 19, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.16522", true),
			Longitude:       null.NewString("106.78182", true),
			Title:           null.NewString("Webinar INTERNET MARKETING REVOLUTION", true),
			Description:     null.NewString("INTERNET MARKETING REVOLUTION Temukan 3 Rahasia SAKTI Bagaimana Memulai Bisnis di Era NEW NORMAL Tanpa Perlu Modal Besar Bahkan Langsung Bisa Anda Praktekkan Setelah Webinar! Anda akan Belajar langsung dari Praktisi Digital Marketing Sejak 2010 dan Sudah Membantu Ribuan Newbie Untuk Memiliki Bisnis Digital TEMUKAN SOLUSINYA dalam Gratis Webinar 90 menit:*  Temukan 3 RAHASIA DALAM WEBINAR ini sbb : RAHASIA #1 Temukan 3 Langkah Sederhana Memiliki Bisnis Ideal di Era New Normal RAHASIA #2 Miliki 4 Persyaratan Terpenting Supaya Bisnis Anda Kebal Krisis RAHASIA #3 Rahasia Top Internet Millionaire Dunia yang Bisa Anda Duplikasi Polanya Karena jumlah Tiket Webinar Sangat Terbatas, Amankan Tiket GRATIS Anda sekarang juga :", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"webinar", "internet", "marketing", "revolution", "bisnis"},
			Image:           null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F293461349%2F312123412152%2F1%2Foriginal.20200605-100422?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C200%2C1080%2C540&s=94fe73bbb2fa3bb25ee5de9350241936", true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("SMAN 3 TARUNA ANGKASA JAWA TIMUR Sman3 Taruna Angkasa, Jalan Ring Road Barat, Ngegong, Manguharjo, Mangu Harjo, Madiun Kota, Jawa Timur, Indonesia", true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 5, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 6, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-7.60562", true),
			Longitude:       null.NewString("111.51571", true),
			Title:           null.NewString("CLOSING SISO 40", true),
			Description:     null.NewString("Closing SISO 40 adalah acara tahunan dari OSISSMAGANASA yang berfungsi sebagai bentuk perayaan dan puncak peringatan ulang tahun SMAN 3 Taruna Angkasa Jawa Timur. Dalam acara ini, kami akan mengumpulkan penonton untuk menonton pertunjukan dari ekstrakurikuler sekolah dan bintang tamu terakhir kami, yaitu \"FOR REVENGE.\"", true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"konser", "musik", "konsermusik"},
			Image:           null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1701940501-koPp8pYC0gpuCNOZXaGdKHJu85rwd21v.png?width=1024&quality=90", true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("SMP SANTA THERESIA, Pangkalpinang", true),
			Category:        []string{"Sports", "Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 3, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 3, 28, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 3, 28, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Now(), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-2.14426", true),
			Longitude:       null.NewString("106.09702", true),
			Title:           null.NewString("THERESIA CUP 2024 ", true),
			Description: null.NewString(`Selamat datang di Event Cup Theresia! Sebuah perhelatan prestisius yang mengundang seluruh komunitas sekolah dan pecinta olahraga untuk merayakan semangat persaingan, kebersamaan, dan prestasi di lingkungan pendidikan.
		
		Dalam atmosfer yang penuh semangat, kami dengan bangga menghadirkan serangkaian perlombaan yang menantang dan memacu kemampuan atletik para peserta. Terutama dalam olahraga basket dan voli, acara ini membuka peluang bagi siswa-siswa berbakat untuk menunjukkan keahlian dan semangat kompetitif mereka di lapangan.
		
		Selain itu, kami juga menyajikan serangkaian kegiatan menarik lainnya yang melibatkan beragam bakat, mulai dari lomba seni, pertunjukan musik, hingga ajang kecerdasan dan pengetahuan. Semua ini dirancang untuk menciptakan pengalaman berkesan bagi setiap peserta dan penonton yang hadir.
		
		Mengusung semangat kebersamaan, kejujuran, dan sportivitas, Event Cup Sekolah Theresia bukan hanya sekadar kompetisi, tetapi juga sebuah perayaan keberagaman dan prestasi di kalangan siswa. Kami mengundang Anda untuk bergabung, merayakan prestasi, dan mendukung semangat fair play yang menjadi inti dari acara ini.
		
		Tunjukkan dukungan Anda, rasakan semangat kompetisi yang menyala, dan jadilah bagian dari perayaan luar biasa di Cup Sekolah Notre Dame. Mari bersama-sama menciptakan kenangan tak terlupakan!`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			IsOnline:      null.NewBool(false, true),
			Tags:          []string{"olimpiade", "kompetisi", "olahraga", "seni", "prestisius"},
			Image:         null.NewString("https://ytknews.id/wp-content/uploads/2022/03/WhatsApp-Image-2022-03-12-at-10.51.40.jpeg", true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Location:        null.NewString("Jiexpo Kemayoran - Jakarta Utara", true),
			Category:        []string{"Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 6, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 6, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 7, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 8, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.14588", true),
			Longitude:       null.NewString("106.84579", true),
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
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			IsOnline:      null.NewBool(false, true),
			Tags:          []string{"cosplay", "japan", "bazaar", "bigbangjakartamatsuri", "streetfood"},
			Image:         null.NewString("https://bbo.co.id/images/modules/bigbang/event/234/bbf-jkt-fest-2023-img-gal.png", true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Entertainment & Performance", "Sports"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 8, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 9, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Now(), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.90760", true),
			Longitude:       null.NewString("107.60148", true),
			Title:           null.NewString("Liga Mahasiswa Basketball blibli.com West Java Conference Season 7", true),
			Description:     null.NewString(`Ikuti keseruan Liga Mahasiswa Basketball blibli.com West Java Conference Season 7! Event ini akan berlangsung dari tanggal 8 hingga 15 Juli 2019 di GOR Pajajaran, Bandung. Anda dapat menonton pertandingan secara langsung atau melalui live streaming di YouTube. Dukung tim kampus Anda dan saksikan siapa yang akan menjadi juara musim ini. Jangan lewatkan aksi-aksi seru dari para pemain basket mahasiswa! Catatan: Untuk informasi lebih lanjut dan live streaming, kunjungi www.youtube.com/ligamahasiswa.`, true),
			AverageRating:   null.NewFloat64(0, true),
			IsFinished:      null.NewBool(false, true),
			IsOnline:        null.NewBool(false, true),
			Tags:            []string{"LIMABasketball", "LigaMahasiswa", "AwalMasaDepan", "LIMA", "studentathlete", "Mahasiswa", "basketball", "basket", "college", "indonesia", "bliblidotcom"},
			Location:        null.NewString("GOR Pajajaran, Bandung", true),
			Image:           null.NewString("https://ngobrolbasket.files.wordpress.com/2015/08/basket-lands.jpg", true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 8, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 9, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			CreatedAt:       null.NewTime(time.Now(), true),
			UpdatedAt:       null.NewTime(time.Now(), true),
			Latitude:        null.NewString("-6.14588", true),
			Longitude:       null.NewString("106.84579", true),
			Title:           null.NewString("Festival Musik Rock Jakarta 2024", true),
			Description: null.NewString(`Festival Musik Rock Jakarta 2024: Nikmati Sensasi Rock Terbaik di Akhir Tahun!

			Apakah Anda penggemar musik rock? Jika ya, maka Anda tidak boleh melewatkan Festival Musik Rock Jakarta 2023, yang akan digelar pada tanggal 27-28 Desember 2023 di Gelora Bung Karno, Jakarta. Festival ini akan menjadi pesta musik rock terbesar di Indonesia, yang akan menampilkan lebih dari 50 band rock papan atas dari dalam dan luar negeri.
			
			Anda akan dapat menyaksikan penampilan spektakuler dari band-band legendaris seperti Slank, Superman Is Dead, Burgerkill, Muse, Linkin Park, Metallica, dan Guns N' Roses, yang akan membawakan lagu-lagu hits mereka di atas panggung. Anda juga akan dapat menemukan band-band baru dan berbakat yang siap mengguncang dunia`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			IsOnline:      null.NewBool(false, true),
			Tags:          []string{"rock, festival, musik, konser"},
			Location:      null.NewString("Jiexpo Kemayoran - Jakarta", true),
			Image:         null.NewString("https://awsimages.detik.net.id/community/media/visual/2023/06/02/poster-konser-one-ok-rock-di-jakarta.jpeg?w=1200", true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Art & Culture", "Sports"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 9, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 9, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("Institut Teknologi Nasional Malang", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"porseni", "itn", "semarang", "olahraga", "basket", "futsal", "theater", "badminton", "esport"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/_thumbnail/600x600/img-20200228-wa0000.jpeg", true),
			Latitude:      null.NewString("-7.95802", true),
			Longitude:     null.NewString("112.61234", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Tech Outlook 2024 - Elevating Recruitment in the Cloud Era", true),
			Description: null.NewString(`Seminar Tech Outlook 2024
		
		Tahun 2024 akan menjadi tahun yang penuh ketidakpastian dalam perekrutan talenta. Tekanan yang tinggi di pasar kerja mendorong pemimpin bisnis untuk mencari solusi inovatif seperti perekrutan dengan kecerdasan buatan (AI) atau menggunakan platform komputasi awan yang efisien.
		
		Alibaba Cloud & Jobstreet by SEEK mengundang kamu bergabung di Seminar Tech Outlook 2024 yang akan diisi narasumber : Muhammad Rohibun ( Solution Architect Alibaba Cloud Indonesia )
		
		Acara : Kamis, 14 Des 2023
		Jam : 12.00 – 16.00 WIB
		Tempat : Bandung
		Link RSVP 👉️ https://zfrmz.com/PVNqoQ7n7WBKTQMq9k4N
		
		Kuota terbatas, segera daftarkan diri kamu di acara ini ( FREE ).`, true),
			Location:      null.NewString("Institut Teknologi Bandung, Kab. Bandung", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"ITB", "tech", "seminar", "AI", "komputasiawan", "cloudcomputing"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4486-tech-outlook-2024-elevating-recruitment-in-the-cloud-era.jpeg", true),
			Latitude:      null.NewString("-6.88901", true),
			Longitude:     null.NewString("107.61027", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("International Conference tentang Teknik, Teknologi Informasi, Ilmu Terapan, Perangkat Lunak Komputer, & Jaringan.", true),
			Description: null.NewString(`Proposal untuk presentasi undangan dapat disampaikan ke eiacn@aet-forum.com. Tujuan dari RFAET Forum adalah menyediakan platform bagi peneliti, insinyur, akademisi, serta profesional industri dari seluruh dunia untuk menyajikan hasil penelitian dan aktivitas pengembangan mereka di Bidang Ilmu Terapan dan Teknologi Rekayasa.

			Konferensi ini memberikan kesempatan kepada para peneliti untuk bertukar ide baru dan pengalaman aplikasi secara langsung, membentuk hubungan bisnis atau penelitian, dan menemukan mitra global untuk kolaborasi masa depan. Batas antara teknologi mutakhir dan inovasi revolusioner membentuk batas depan komputasi yang harus didorong untuk menyediakan dukungan yang diperlukan bagi kemajuan lebih lanjut dalam berbagai bidang teknik dan teknologi. Platform bersama ini diharapkan dapat memberikan dasar untuk kemitraan di antara berbagai bidang untuk melayani masyarakat dengan lebih baik.
			
			Semua abstrak, makalah, dan poster yang diajukan akan melalui proses peer review tanpa pengetahuan asal, dan manuskrip yang diterima akan diterbitkan dalam prosiding konferensi. Makalah konferensi yang dipilih akan diterbitkan di jurnal yang terkait dengan konferensi ini. Jurnal terkait akan mengalokasikan isu khusus/teratur untuk makalah yang diajukan ke konferensi ini.
			
			Manfaat Bergabung:
			Sering kali diamati bahwa orang kurang memiliki motivasi dan kepercayaan diri untuk berpartisipasi dalam acara internasional, terutama karena hambatan yang dibuat sendiri atau budaya. Kami mengakui faktor ini dan memastikan mendukung dan memotivasi mereka yang baru pertama kali dan juga akademisi berpengalaman dengan:
			
			Membangun hubungan akademis dan profesional Anda
			Meningkatkan semangat dan kepercayaan diri Anda dalam menyajikan penelitian Anda di platform internasional
			Mengatasi hambatan Anda untuk beradaptasi dengan lingkungan asing
			Menyediakan pengalaman holistik dari pariwisata akademis`, true),
			Location:      null.NewString("Pullman Jakarta, Central Park", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"konferensi", "edukasi", "conference", "teknologiinformasi", "software"},
			Image:         null.NewString("https://ocs.teknokrat.ac.id/public/conferences/2/homepageImage_en_US.png", true),
			Latitude:      null.NewString("-6.17800", true),
			Longitude:     null.NewString("106.78941", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(2, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 14, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("Centennial Tower 36th Floor, Jl.Jend Gatot Subroto Kav. 24-25, Jakarta Indonesia Jakarta Selatan, Jakarta 12930", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"DPO", "edukasi", "data", "course", "pelatihan", "kursus"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F644815509%2F572826263471%2F1%2Foriginal.20231120-024306?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C2160%2C1080&s=b05c90f991721ed07c8626f3ff5634e4", true),
			Latitude:      null.NewString("-6.23059", true),
			Longitude:     null.NewString("106.82106", true),
			GuestStar:     []string{"Prof. Dr. IBR Supancana", "Prof. Abu Bakar Munir"},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("JOB FAIR BERSAMA PUSAT KARIR PERGURUAN TINGGI SOLORAYA 2023", true),
			Description: null.NewString(`JOBFAIR AKBAR, TERBESAR, DAN TERLENGKAP DI SOLORAYA
		
		DI USUNG LEBIH DARI 20 KAMPUS TERNAMA
		
		DIIKUTI PULUHAN PERUSAHAAN DAN RIBUAN LOWONGAN PEKERJAAN
		
		
		
		JOB FAIR BERSAMA
		
		PUSAT KARIR PERGURAN TINGGI
		
		SOLORAYA 2023
		
		Acara ini didukung oleh kampus merdeka yang bertujuan untuk membantu mahasiswa dan alumni Universitas dalam mencari pekerjaan atau magang. Terdapat lebih dari 50 perusahaan yang akan berpartisipasi dalam acara ini, termasuk perusahaan-perusahaan ternama seperti Microsoft, Google, dan Amazon. Selain itu, acara ini juga akan diisi dengan seminar dan workshop yang akan membahas berbagai topik menarik seputar dunia kerja, seperti cara membuat CV yang menarik, tips wawancara kerja, dan banyak lagi.`, true),
			Location:      null.NewString("Gedung Grha Soloraya (Eks Gedung Bakorwil) Surakarta, Kota Surakarta", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"jobfair", "edukasi", "kampusmerdeka"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4438-job-fair-bersama-pusat-karir-perguran-tinggi-soloraya-2023.jpeg", true),
			Latitude:      null.NewString("-7.57276", true),
			Longitude:     null.NewString("110.82880", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 13, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 17, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Charity Children Camp 2023 - SERENITY", true),
			Description: null.NewString(`🪐 FKIK Atma Jaya Presents: Charity Children Camp 2023 - SERENITY 🪐
			Sesuai dengan nama serenity, mahasiswa FKIK Atma Jaya berharap untuk bisa memberikan kebahagiaan, ketenangan, dan kedamaian untuk pesertanya. 
			Charity Children Camp tahun ini memiliki tujuan untuk menyediakan sarana edukasi dan dukungan psikososial yang diharapkan dapat memotivasi dan membangkitkan semangat hidup anak-anak penderita Down Syndrome. Selain itu, dengan adanya acara ini diharapkan dapat meruntuhkan stigma masyarakat mengenai Down Syndrome dan mengenal mereka lebih baik dengan cara berinteraksi serta bermain bersama mereka. Acara ini akan dilaksanakan pada 18-19 November, tertutup untuk mahasiswa FKIK Atma Jaya.
			Akan tetapi, Charity Children Camp memberikan kesempatan untuk masyarakat dalam mendukung anak-anak Down Syndrome dari Rumah Ceria Down Syndrome - POTADS sebagai Non-Governmental Organization yang bekerja sama dengen FKIK Atma Jaya, dengan cara berdonasi. Jika saudara/i ingin berdonasi, link dapat diakses melalui http://bit.ly/3FQd59t atau pada barcode yang tersedia pada poster. 
			“We have the same wants and dreams as everyone else. We can do anything anyone else can do. We are more alike than we are different.” says Manager of Advocacy at the National Down Syndrome Society`, true),
			Location:      null.NewString("Wisma Kinasih, Depok", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"charity", "downsyndrome", "ayoberbagi", "FKIKAtma", "atmajaya", "donasi"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/22/4480-charity-children-camp-2023-serenity.jpeg", true),
			Latitude:      null.NewString("-6.45367", true),
			Longitude:     null.NewString("106.86311", true),
			GuestStar:     []string{"Dr. dr. Bina Akura Sp.A(K)"},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 17, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 6, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 6, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("2024 6th International Conference on Intelligent Medicine and Image Processing (IMIP 2024)", true),
			Description: null.NewString(`SELAMAT DATANG DI IMIP 2024!
			Konferensi Internasional ke-6 tentang Kedokteran Cerdas dan Pengolahan Citra (IMIP 2024) akan berlangsung di Bali, Indonesia pada tanggal 18 Juni 2024. IMIP 2024 disponsori oleh Universitas Udayana, bersama-sama dengan Universitas Tiangong dan Masyarakat Biologi dan Bioinformatika.
			
			Kedokteran cerdas dan pengolahan citra telah mencapai kemajuan luar biasa dalam domain medis, membawa inovasi yang transformasional dalam diagnosis dan pengobatan. Pembelajaran mesin dan kecerdasan buatan memainkan peran kunci dengan memeriksa gambar medis, menyusun solusi medis personal, memungkinkan pemantauan jarak jauh, dan memfasilitasi pengembangan obat. Secara bersamaan, teknologi pengolahan citra meningkatkan kualitas gambar medis, memfasilitasi diagnosis yang lebih akurat. Inovasi ini tidak hanya meningkatkan efisiensi dan efektivitas biaya medis tetapi juga memberikan kemampuan deteksi penyakit lebih awal, rencana perawatan personal, dan pengalaman medis yang lebih baik secara keseluruhan kepada pasien.
			
			Seri konferensi IMIP telah sukses diadakan di lokasi seperti Bali, Indonesia, dan Tianjin, China. Sebagai acara tahunan, IMIP berusaha untuk membentuk platform komunikasi bagi pemimpin teknologi terkemuka, sarjana, insinyur, ilmuwan, ahli industri, dan mahasiswa pascasarjana untuk bertukar ide dan menjelajahi perkembangan teknologi terbaru dalam Kedokteran Cerdas dan Pengolahan Citra, atau bidang terkait. Konferensi ini melibatkan pidato utama, pembicara tamu, sesi presentasi lisan dan poster, presentasi video, pameran, dan format interaktif lainnya, memfasilitasi komunikasi dan pertukaran informasi yang kuat dalam bidang tersebut.
			Kami berharap untuk bertemu Anda di Bali yang indah, Indonesia!`, true),
			Location:      null.NewString("Bali, Indonesia", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"konferensi", "edukasi", "bioinformatika", "informatika", "bioinformatics", "imageprocessing", "medicine"},
			Image:         null.NewString("https://www.imip.org/images/slide2.jpg", true),
			Latitude:      null.NewString("-8.40402", true),
			Longitude:     null.NewString("115.19064", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("PROGRAM ASISTENSI SEKOLAH: OPEN DONASI", true),
			Description: null.NewString(`‼PROGRAM ASISTENSI SEKOLAH: OPEN DONASI ‼ Halo, Sobat PAS! 🙌🏻
			Kami dari Program Asistensi Sekolah (PAS) dengan senang hati mengajak teman-teman semua untuk ikut berpartisipasi dalam donasi untuk kegiatan kami yang bertujuan membantu anak-anak di SDN Kebon Kopi Bogor dan SDN Pancoranmas 3 Depok menjadi lebih baik. Kebaikan dalam bentuk apapun akan sangat berarti untuk mereka yang ada di sekolah😊
			DONASI BARANG: 
			📦Dikumpulkan dalam wadah 📍Titik Penyerahan: FIB UI atau Asrama UI
			☎: Yelena (082246266727)
			DONASI UANG: 
			💰Transfer via Bank BNI (1230498153) a.n. Najwa Salsabila
			Mari bersama-sama memberikan harapan dan masa depan yang lebih baik untuk anak-anak pada sekolah dasar. Setiap kontribusi Anda sangat berarti bagi mereka.`, true),
			Location:      null.NewString("FIB UI, Kota Depok", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"charity", "downsyndrome", "ayoberbagi", "FKIKAtma", "atmajaya", "donasi"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4469-program-asistensi-sekolah-open-donasi.jpeg", true),
			Latitude:      null.NewString("-6.36343", true),
			Longitude:     null.NewString("106.82856", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Sports"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 5, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("FUN WALK 90's", true),
			Description: null.NewString(`Fun Walk 90’s
			Local Brand Fashion Market & Kuliner
			Music Performance
			Result
			KESAN, Keseruan yang berbeda di tengah kebosanan gelaran event olahraga santai
			EXPOSURE, Digital Traffic meningkat karena peserta Upload keseruan saat Event Berlangsung
			TRUST, Kepercayaan masyarakat dan sponsor terhadap STIE TOTALWIN meningkat
			Dengan target minimal 5000 peserta akan mengikuti FUN WALK 90's yang digelar oleh STIE TOTALWIN. Adapun ragam hiburan selain keseruan fashion para peserta diantaranya Local Brand Fashion Market 90an dan puluhan kuliner & jajanan sepanjang Jalan Suratmo, Kota Semarang.`, true),
			Location:      null.NewString("Kampus STIE Totalwin, Kota Semarang", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"olahraga", "kuliner", "jalan-jalan", "hiburan"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4454-fun-walk-90s_1.jpeg", true),
			Latitude:      null.NewString("-7.00080", true),
			Longitude:     null.NewString("110.38800", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 11, 10, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 11, 17, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("FIB UI, Kota Depok", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"esport", "dota2", "mobilelegends", "PUBG", "Valorant", "esportUI"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/_thumbnail/600x600/ui-battlegrounds-2021-feeds-ig.jpg", true),
			Latitude:      null.NewString("-6.35691", true),
			Longitude:     null.NewString("106.82712", true),
			GuestStar:     []string{"Anisa Rahim - Onic Esport"},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("Persada Indonesia University, Jakarta Pusat", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"businesscase", "garena", "lomba", "kompetisi"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/_thumbnail/600x600/p.jpg", true),
			Latitude:      null.NewString("-6.19538", true),
			Longitude:     null.NewString("106.84874", true),
			GuestStar:     []string{"Garena Indonesia"},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("UNIVERSITAS WIDYATAMA. Jakarta Pusat", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"esport", "mobilelegends", "PUBG", "kompetisi"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/pamflet-story.jpg", true),
			Latitude:      null.NewString("-6.89776", true),
			Longitude:     null.NewString("107.64542", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 12, 4, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 12, 10, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 12, 20, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("SMA NEGERI 52 JAKARTA", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"esport", "olahraga", "Smandu cup", "futsal", "mobile legends", "basket", "badminton", "volley"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2020/01/MP-Smaducup-2020-Robhustrophica-SMAN-52-Jakarta-Copy.jpg", true),
			Latitude:      null.NewString("-6.12224", true),
			Longitude:     null.NewString("106.92410", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("Universitas YARSI, Jakarta Pusat", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"menulis", "essay", "opini", "kompetisi"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/_thumbnail/600x600/whatsapp-image-2021-04-28-at-114907.jpg", true),
			Latitude:      null.NewString("-6.16949", true),
			Longitude:     null.NewString("106.87014", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Art & Culture", "Charity", "Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("Cinepolis Lippo Plaza Yogyakarta", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"teater", "pentas", "ugm", "charity"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/21/_thumbnail/600x600/4329-charity-and-cultural-night-amsa-ugm-2022.jpeg", true),
			Latitude:      null.NewString("-7.78390", true),
			Longitude:     null.NewString("110.39077", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 18, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("OPENING GEBYAR NUSANTARA 2021", true),
			Description: null.NewString(`Salam Budaya!

			Apa sih Gebyar Nusantara itu?
			Gebyar Nusantara (@gebyarnusantaraipb) merupakan rangkaian acara kebudayaan terbesar di IPB University yang dilaksanakan oleh Kementrian Seni dan Budaya BEM KM IPB
			
			Gebyar Nusantara 2021 mengangkat tema “Abhipraya Palawa Nusantara” yang memiliki arti harapan nusantara yang bersemi kembali. Gebyar Nusantara diikuti oleh 26 Organisasi Mahasiswa Daerah dan akan dimeriahkan dengan berbagai rangkaian acara mulai dari Opening, Bincang Budaya, Pemilihan Putera-Puteri OMDA, Video Creative Competition, dan Malam Puncak.
			Saksikan kemeriahannya dengan special performance Kisah Abhipraya dan mari bersama turut lestarikan warisan budaya nusantara!
			`, true),
			Location:      null.NewString("Kota Bogor", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"nusantara", "budaya", "seni"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/21/_thumbnail/600x600/4329-charity-and-cultural-night-amsa-ugm-2022.jpeg", true),
			Latitude:      null.NewString("-6.59792", true),
			Longitude:     null.NewString("106.80150", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Art & Culture", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 18, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 18, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 18, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("UNNES, Semarang", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"nusantara", "budaya", "seni", "makeup", "kompetisi", "talentshow"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/whatsapp-image-2021-06-04-at-32244-pm.jpg", true),
			Latitude:      null.NewString("-7.05048", true),
			Longitude:     null.NewString("110.39243", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 11, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 11, 18, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Agrifest 4.0", true),
			Description: null.NewString(`Agrifest merupakan festival musik tahunan yang selalu diadakan di Fakultas Pertanian , Universitas Sumatera Utara. Nama Agrifest merupakan festival musik ke-4, setelah sebelumnya ada beberapa nama lain. Agrifest merupakan acara puncak perayaan Ulang Tahun IMASEP Ke-39 dan Agribisnis Ke-59. Sebelumnya ada beberapa kegiatan yang dilaksanakan seperti : pengabdian masyarakat yang diadakan 5 Oktober 2019, kemudian dilanjutkan Agribusiness Competition yang dimeriahkan oleh semua mahasiswa agribisnis untuk tujuan mempererat tali persaudaraan lagi. Kemudian ada kegiatan perayaan Ulang Tahun IMASEP Ke-39 dan Agribisnis Ke-59. Dan acara terakhirnya yaitu AGRIFEST. 
Event Agrifest selalu mengundang artis ibukota dan dimeriahkan oleh band-band lokal medan, seperti di Tahun 2017 kemarin mengundang Fourtwnty, di tahun 2018 mengundang Fiersa Besari x Kerabat Kerja, dan tahun ini mengundang Senar Senja, Dimana lagu Senar Senja yang sedang booming yaitu "Dialog Hujan", "Menua Berdua", "Havanah", "Untuk Yang Baru Saja di Wisuda", "Malu tapi Mau", "Asmara Roda Dua", dan masih banyak lagi. Tidak hanya menampilkan artis saja, juga ada band lokal yang ikut memeriahkan seperti : Hello Benji, The Sugarcane, The Boxquitos, Pesawat Sederhana, dan Not Xmprewell. Ada festival makanan juga yang ikut memeriahkan. Semua Makanan Dijamin Enakkkkk!!!!! 
			`, true),
			Location:      null.NewString("Fakultas Pertanian USU, Kota Medan", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"agrifest", "musik", "konser"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/2/_thumbnail/600x600/agrifest-40.jpeg", true),
			Latitude:      null.NewString("3.55713", true),
			Longitude:     null.NewString("98.65468", true),
			GuestStar:     []string{"Senar Senja", "TheBoxQuitos", "The Sugar Cane", "Pesawat Sederhana", "Not Xmprewell"},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 18, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString("Fakultas Ilmu Budaya Universitas Airlangga, Kota Surabaya", true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"agrifest", "musik", "konser"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/3/malam-puisi-airlangga.jpeg", true),
			Latitude:      null.NewString("-7.27287", true),
			Longitude:     null.NewString("112.75987", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 18, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC), true),
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
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"jepang", "musik", "konser", "japan", "cosplay", "fashion", "anisong"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702308316-mzPIzZcALZDPlkHzHGHDjNgRQg1nCB5x.jpg?width=1024&quality=90", true),
			Latitude:      null.NewString("-6.21203", true),
			Longitude:     null.NewString("106.80533", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 14, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 20, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 21, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("BAHTERA VOL. 02", true),
			Description: null.NewString(`Bahtera Vol.2 merupakan event lanjutan dari Bahtera Vol.1 yang dimana HangtuahCup akan kembali mengadakan Pensi Festival Musik sebagai penutupan acara pada tanggal 27 november 2023, event pensi yang diadakan akan mengundang beberapa guest star untuk mengisi event tersebut, pada pensi ini merupakan event yang dijalani oleh siswa siswi SMA HANGTUAH 1 JAKARTA sebagai bentuk memeriahkan hangtuah cup 2024!
			`, true),
			Location: null.NewString(`GOR Seskoal
			RT.1/RW.11, Cipulir, Kby. Lama, Kebayoran Lama, Jakarta Selatan, Jakarta, Indonesia
			GOR SESKOAL`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"konser", "musik", "bahtera", "hangtuah"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1704444134-5FIrt3yyNk7N1QI98jgXw9jklLN0WTxV.PNG?width=1024&quality=90", true),
			Latitude:      null.NewString("-6.23931", true),
			Longitude:     null.NewString("106.77159", true),
			GuestStar:     []string{"Nadin Anizah", "Nadhif Basalamah", "Yahya"},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 20, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 22, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 6, 24, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 6, 24, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Central Park, Jakarta Barat`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"seminar", "workshop", "tidur", "sleep"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1703650368-Eaiba4iQyVkSOiq5lgcWXoTkm6rOLEB1.png?width=1024&quality=90", true),
			Latitude:      null.NewString("-6.17642", true),
			Longitude:     null.NewString("106.79120", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 25, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 26, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 7, 2, 0, 0, 0, 0, time.UTC), true),
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
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"talkshow", "workshop", "crypto"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1703650368-Eaiba4iQyVkSOiq5lgcWXoTkm6rOLEB1.png?width=1024&quality=90", true),
			Latitude:      null.NewString("-6.22144", true),
			Longitude:     null.NewString("106.65420", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career", "Art & Culture", "Entertainment & Performance", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 25, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 26, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Nanotenri 2023", true),
			Description: null.NewString(`Festival budaya Jepang yang diselenggarakan oleh ekskul Bahasa Jepang SMA Labschool Cibubur dengan lomba - lomba seperti cosplay, obakeyashiki, hanabi, special performance, karaoke, menggambar manga dan lain-lain. 

			Menampilkan berbagai performance, dari peserta lomba hingga guest star spesial kami! 
			`, true),
			Location: null.NewString(`SD - SMP - SMA LABSCHOOL CIBUBUR
			Jl. Raya Hankam Kampus Labschool No.15-20, Jatiranggon, Kecamatan Jatisampurna, Bekasi Kota, Jawa Barat, Indonesia`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"jepang", "cosplay", "labschool"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1697986623-dgOWDEbq7mkzMlNH3imPFRfx8xcCO8SI.png?width=1024&quality=90", true),
			Latitude:      null.NewString("-6.34319", true),
			Longitude:     null.NewString("106.92521", true),
			GuestStar:     []string{"Elaine"},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity", "Sports"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 25, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 26, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Jl. Meruya Selatan, No. 19, Kembangan - Jakarta Barat`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"donor", "donor darah", "senam sehat", "proyekrumah", "renovasirumahrenotop"},
			Image:         null.NewString("https://renotop.id/img/blogs/donor-darah-di-renotop-meruya-17-juni-2023-Slsfq.jpg", true),
			Latitude:      null.NewString("-6.21173", true),
			Longitude:     null.NewString("106.73739", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 28, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Jungleland Sentul`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"cosplay", "musik", "konser", "JKT48", "akhirtahun"},
			Image:         null.NewString("https://images.genpi.co/uploads/arsip/normal/2022/11/30/ilustrasi-menikmati-momen-libur-akhir-tahun-di-c-tfur.jpg", true),
			Latitude:      null.NewString("-6.57341", true),
			Longitude:     null.NewString("106.89222", true),
			GuestStar:     []string{"JKT48"},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Aeon Mall BSD City`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"makeup", "expo", "aeon", "makeupexpo", "BSD"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2022/08/PP-EJ-Indonesia-Makeup-Expo-IMAE-@imaeofficial-Copy.jpg", true),
			Latitude:      null.NewString("-6.30414", true),
			Longitude:     null.NewString("106.64367", true),
			GuestStar:     []string{"JKT48"},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Online`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			IsOnline:      null.NewBool(true, true),
			Tags:          []string{"basket", "kompetisi", "freethrow"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2020/10/MP-EJ-Permata-Freethrow-Championship-2020-Mahasiswa-Unsoed-UNS-dan-UPI-Copy.jpg", true),
			Latitude:      null.NewString("0.0", true),
			Longitude:     null.NewString("0.0", true),
			GuestStar:     []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo", "Entertainment & Performance", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 28, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Kota Tua, Glodok Plaza`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"festival", "kotatua", "kuliner", "cosplay", "livemusic", "musik"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2023/12/PP-EJ-Festival-Glodok-Kuliner-Kota-Tua-JML-Organizer-85-Copy.jpg", true),
			Latitude:      null.NewString("-6.14208", true),
			Longitude:     null.NewString("106.81654", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 30, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 30, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Digital Lounge CIMB Niaga, Central Park`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"UI", "UX", "Design Thinking", "Seminar", "Workshop", "UI/UX"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2023/12/MP-EJ-UpSkill-1.0-KEJAR-MIMPI-TANGSEL-1-Copy.jpg", true),
			Latitude:      null.NewString("-6.17781", true),
			Longitude:     null.NewString("106.78947", true),
			GuestStar:     []string{"Aswin Widyastama", "Miqdad Darmawan"},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 9, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("2024 the 7th International Conference on Big Data and Smart Computing (ICBDSC 2024)", true),
			Description: null.NewString(`Panitia Penyelenggara dengan senang hati mengumumkan bahwa Konferensi Internasional ke-7 tentang Big Data dan Smart Computing (ICBDSC 2024) akan diselenggarakan secara daring di Jakarta pada tanggal 23-25 Januari 2024.

			Big Data muncul dalam beberapa tahun terakhir sebagai paradigma baru yang menyediakan data dan peluang melimpah untuk meningkatkan dan/atau memungkinkan aplikasi penelitian dan dukungan keputusan dengan nilai luar biasa untuk aplikasi bumi digital termasuk bisnis, ilmu pengetahuan, dan rekayasa. Smart computing juga merupakan topik hangat yang menggabungkan teknologi komputer canggih untuk menciptakan sistem, aplikasi, dan layanan baru dalam berbagai aplikasi seperti bisnis, perawatan kesehatan, sistem industri, dan sebagainya.
			
			Big data dan smart computing adalah bidang penelitian yang sedang berkembang yang telah menarik perhatian besar dari bidang ilmu komputer, komunikasi dan kontrol, teknologi informasi, serta ilmu sosial dan disiplin lainnya. Dengan volume, kecepatan, dan jenis data besar yang terus berkembang dari sistem jaringan, smart computing menjadi sangat penting untuk menjamin fungsionalitas kritis dalam sistem jaringan, seperti kesadaran situasional real-time di area luas, manajemen data dinamis, optimasi dan kontrol efisiensi, kinerja jaringan yang tangguh, dan sebagainya.
			
			Konferensi ini ditujukan untuk akademisi, peneliti, dan profesional dengan minat khusus terkait topik konferensi. Ini mengumpulkan akademisi, peneliti, dan profesional di bidang Big Data dan Smart Computing, menjadikan konferensi sebagai platform yang sempurna untuk berbagi pengalaman, memperkuat kolaborasi antara industri dan akademisi, dan mengevaluasi teknologi yang sedang berkembang di seluruh dunia.
			`, true),
			Location:      null.NewString(`Jakarta Selatan Kuningan, Lotte Avenue Mall`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"konferensi", "bigdata", "AI", "smartcomputing", "penelitian", "sains", "science", "engineering", "edukasi"},
			Image:         null.NewString("https://www.icbdsc.org/style/images/venue/1.jpg", true),
			Latitude:      null.NewString("-6.22403", true),
			Longitude:     null.NewString("106.82325", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("2024 4th Asia Conference on Information Engineering (ACIE 2024)", true),
			Description: null.NewString(`Konferensi ACIE bertujuan untuk mengumpulkan ilmuwan akademis terkemuka, profesor, 
			peneliti, mahasiswa, dan sarjana penelitian untuk bertukar dan berbagi pengalaman serta hasil penelitian mereka di semua aspek 
			Information Engineering. Ini juga memberikan platform interdisipliner utama bagi peneliti, praktisi, dan pendidik untuk menyajikan 
			dan mendiskusikan inovasi, tren, serta tantangan praktis yang dihadapi dan solusi yang diadopsi dalam bidang Information Engineering.
			`, true),
			Location:      null.NewString(`TerasKota, BSD`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"engineering", "informationengineering", "IT", "konferensi", "conference", "peneliti", "edukasi"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F539205879%2F293320802810%2F1%2Foriginal.20230620-084417?w=1000&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C22%2C542%2C271&s=66d9d4d3e7e12e49093eeabc2ac39574", true),
			Latitude:      null.NewString("-6.29541", true),
			Longitude:     null.NewString("106.66749", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Charity"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Mari Bantu Moestopo Jelajah Nusantara 2024", true),
			Description: null.NewString(`Mari Bantu Moestopo Jelajah Nusantara 2024

			Deskripsi Event :
			❗#BantuMJN2024 ❗
			
			Menurut data dari Riset Kesehatan Dasar (Riskesdas) tahun 2018, prevalensi nasional masalah gigi dan mulut di Indonesia adalah 57,6 persen dan hanya 10,2 persen yang telah mendapatkan pelayanan dari tenaga medis. Menurut laporan Riskesdas Provinsi Maluku Tahun 2018, Provinsi Maluku termasuk ke dalam daerah 3 tertinggi dengan angka prevalensi gigi rusak/berlubang/sakit yaitu sebesar 56,28%.
			
			Yuk bantu kami untuk terwujudnya Senyum Sehat, Senyum Ceria Masyarakat Maluku melalui link berikut: https://kitabisa.com/campaign/yukbantumjn24
			
			Bantuan dari teman-teman sangat berarti untuk terwujudnya Senyum Sehat, Senyum Ceria Masyarakat Maluku #MengabdiDenganHatiBergerakDenganAksi #BantuMJN2024.
			`, true),
			Location:      null.NewString(`Online`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			IsOnline:      null.NewBool(true, true),
			Tags:          []string{"charity", "kesehatan", "mulut", "gigi", "Maluku"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2023/12/MP-EJ-Moestopo-Jelajah-Nusantara-2024-Senat-Mahasiswa-FKG-Senat-Mahasiswa.jpg", true),
			Latitude:      null.NewString("0.0", true),
			Longitude:     null.NewString("0.0", true),
			GuestStar:     []string{""},
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 11, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 11, 13, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 11, 15, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Universitas TAMA JAGAKARSA, JALAN TB Simatupang No. 152 Jagakarsa, DKI Jakarta`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"sport", "olahraga", "futsal", "lomba"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2023/02/MP-EJ-STUDENT-CUP-UKM-FUTSAL-TAMA-JAGAKARSA-Copy.jpg", true),
			Latitude:      null.NewString("-6.30236", true),
			Longitude:     null.NewString("106.84097", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 11, 4, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`Wahana Ekspresi Pusponegoro`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"sport", "olahraga", "futsal", "lomba", "Basket", "Kompetisi"},
			Image:         null.NewString("https://eventsurabaya.net/wp-content/uploads/2022/06/MP-ES-SNBC-2022-MPK-OSIS-SMAN-1-GRESIK-Copy.jpg", true),
			Latitude:      null.NewString("-7.16190", true),
			Longitude:     null.NewString("112.65370", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 12, 13, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("INFINITUM CUP Badminton Competition", true),
			Description: null.NewString(`INFINITUM CUP Badminton Competition
			Tanggal : 17 Desember 2022
			Tempat : GOR Cempaka Putih, Jakarta Pusat
			
			Deskripsi Event :
			Hai
			
			Kabar baik nih untuk temen-temen yang masih di jenjang SMA. Kami akan melakukan event Badminton nihh, pada tanggal 17 Desember 202 di Gor Cempaka Putih. Pastinya untuk kalian yang mau ikutan boleh langsung regis melalui link diatas yaa.
			
			Oh iya jangan lupa menangin hadiah nya, ya gais. See you!!!!
			`, true),
			Location:      null.NewString(`GOR Cempaka Putih, Jakarta Pusat`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"sport", "olahraga", "badminton", "lomba", "infinitum", "Kompetisi"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2022/11/MP-EJ-Lomba-Badminton-Infinitum-Copy.jpg", true),
			Latitude:      null.NewString("-6.17029", true),
			Longitude:     null.NewString("106.87290", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Art & Culture", "Entertainment & Performance", "Competition"},
			PreregisterDate: null.NewTime(time.Date(2024, 11, 20, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 11, 21, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 12, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Valskrie Korean Journey with Ilkom UBakrie", true),
			Description: null.NewString(`Prodi S1 Ilmu Komunikasi UBakrie mempersembahkan Event Bergengsi “Valskrie”, dan kali ini Budaya yang ditampilkan adalah Budaya Korea Selatan dengan tema “Korean Journey with Ilkom UBakrie”.

			Akan banyak penampilan menarik, unik, seru yang akan menggemparkan panggung Ilkom Universitas Bakrie.
			
			Ikuti lomba Dance Cover dan Song Cover sesuai petunjuk pada slide kedua. Menangkan hadiah jutaan rupiah dan para Pemenang akan tampil di puncak acara Valskrie pada 23 Desember 2023 di Universitas Bakrie.
			`, true),
			Location:      null.NewString(`Universitas Bakrie, Jakarta`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"sport", "dance", "song", "cover", "korean", "korea", "kpop"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2023/11/MP-EJ-Valskrie-Korean-Journey-with-Ilkom-UBakrie.jpg", true),
			Latitude:      null.NewString("-6.22014", true),
			Longitude:     null.NewString("106.83324", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},
		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo", "Sports", "Art & Culture"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 21, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
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
			Location:      null.NewString(`UMN, Gading Serpong`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"sport", "olahraga", "budaya", "esport", "maxima", "UMN", "Gading Serpong"},
			Image:         null.NewString("https://eventjakarta.com/wp-content/uploads/2023/09/MP-EJ-MAXIMA-2023-UMN-1-Copy.jpg", true),
			Latitude:      null.NewString("-6.25601", true),
			Longitude:     null.NewString("106.61858", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 19, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("2024 6th Asia Pacific Information Technology Conference (APIT 2024)", true),
			Description: null.NewString(`Welcome to APIT 2024!
The 2024 6th Asia Pacific Information Technology Conference (APIT 2024) will be held during 29-31 January, 2024 in Bangkok, Thailand. APIT 2024 is an international forum for sharing knowledge and results in theory, methodology and new advances and research results in the fields of Information Technology. The conference will bring together researchers and practitioners from both academia as well as industry to meet and share cutting-edge development in the field. The conference welcomes significant contributions in all major fields of the Information Technology in theoretical and practical aspects. It will put 
special emphasis on the participations of PhD students, Postdoctoral fellows and other young researchers from all over the world. It would be beneficial to bring together a group of experts from diverse fields to discuss recent progress and to share ideas on open questions. The conference will feature world-class keynote speakers in the main areas.
			
			Humans have been storing, retrieving, manipulating, and communicating information since the  Sumerians  in  Mesopotamia  developed writing in about 3000 BC, but the term information technology in its modern sense first appeared in a 1958 article published in the Harvard Business Review; authors Harold J. Leavitt and Thomas L. Whisler commented that "the new technology does not yet have a single established name. We shall call it information technology (IT)." Their definition consists of three categories: techniques for processing, the application of statistical and mathematical methods to decision-making, and the simulation of higher-order thinking through computer programs. What is the future of IT, and which technology is going to the rule the IT industry?  Let us join in APIT 2024 to enjoy the discussion!
			`, true),
			Location:      null.NewString(`The Cartensz Mall Main Atrium, Gading Serpong`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"conference", "konferensi", "informationtechnology", "technology"},
			Image:         null.NewString("https://www.apit.net/style/images/banner-4.jpg", true),
			Latitude:      null.NewString("-6.26975", true),
			Longitude:     null.NewString("106.63059", true),
			GuestStar:     []string{""},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 19, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Swastamita.fest", true),
			Description: null.NewString(`Swastamita.fest merupakan sebuah event konser musik yang Bertemakan Let It Flow With Swastamitafest, Guest star yang akan dibawakan ini bergenre pop dan indie yang akan dilaksanakan di kota bandung, Musik yang diputar dari seniman terkenal akan hadir pada malam yang tak akan terlupakan dari Event musik dan hiburan live di SwastamitaFest !!!

			Dengan Phase 1
			1. The Panturas !!!
			2. Aruma !!!
			3. Dikta Wicaksono !!!
			
			Yuk kira-kira bakal ada siapa lagi Guest starnya.
			
			Oh iya selain itu bakal ada band lokal juga yang bakal hadir, pokoknya Jangan lewatkan kesempatan ini untuk menciptakan kenangan indah dengan teman, Pasangan kalau ada dan keluarga. Beli tiket Anda sekarang untuk malam kegembiraan !
			`, true),
			Location:      null.NewString(`Lapangan PPI Pussenif Jl. Cisadea No.4, Bandung Kota, Bandung Kota, Jawa Barat, Indonesia`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"musik", "konser", "Dikta Wicaksono", "Bandung"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1706762357-o2PeYcWKyozxdD5SclqI2koMZT1T3Bm1.jpg?width=1024&quality=90", true),
			Latitude:      null.NewString("-6.90552", true),
			Longitude:     null.NewString("107.63321", true),
			GuestStar:     []string{"The Panturas", "Aruma", "Dikta Wicaksono"},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Art & Culture", "Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 19, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Connx Off The Weekend (OTW) Festival", true),
			Description: null.NewString(`Connx "Off the Weekend" / OTW adalah festival besar di akhir pekan! Di festival ini, kamu bisa seru-seruan dengan musisi lokal dan juga internasional. Nggak cuma itu, festival ini juga merangkul beragam komunitas lokal dengan beragam kegiatan seru, seperti konser musik, pameran seni, dan workshop. Jadi, ayo habiskan waktu akhir pekan kamu dengan teman-teman di sini, sambil menikmati suasana yang menyenangkan dan menginspirasi.

			Dengan berbagai artis/penyanyi Phase 1 :
			1. Atlesta feat Dux Stella Voce Choir
			2. Monohero
			3. MRLD (Artist from Philippines)
			4. Nadin Amizah
			5. Salma Salsabil
			6. Shaggydog feat Malang Youth Orchestra
			7. Wake up, Iris
			
			Dan jangan sampai terlewatkan karena masih ada penyanyi lainnya di Phase berikutnya loh! Beli tiketnya sekarang dan nikmati festival kali ini bersama orang tersayang. 
			`, true),
			Location: null.NewString(`Lapangan Rampal
			Jalan Urip Sumoharjo, Klojen, Malang Kota, Jawa Timur, Indonesia`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"musik", "konser", "Seni", "Jatim", "Salma Salsabil"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1708756824-c6rdDAxdiaGd2DHoHb8O8CQ6eGAiDXtA.jpg?width=1024&quality=90", true),
			Latitude:      null.NewString("-7.97356", true),
			Longitude:     null.NewString("112.64042", true),
			GuestStar:     []string{"Salma Salsabil", "Nadin Amizah", "Atlesta feat Dux Stella Voce Choir"},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 19, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("RAWS CHARITY RUN 2024", true),
			Description: null.NewString(`RAWS CHARITY RUN 2024 yang akan diselenggarakan pada 18 Februari 2024 di ECCOS Living Plaza Bali dengan konsep kegiatan Charity Fun Run 11K, bertujuan untuk mempromosikan wisata Bali dalam konsep wisata Kesehatan dengan menerapkan konsep health, wellness dan lifestyle. Event ini akan dilaksanakan di Bali peserta terdiri dari pengusaha dan masyarakat umum. Kegiatan RAWS CHARITY RUN 2024 ini diadakan bersamaan dengan Bali Health & Sport Tourism Expo atau BHESTE yang akan dilaksanakan pada tanggal 16-18 Februari 2024 di Eccos Living Plaza Bali.
			`, true),
			Location: null.NewString(`ECCOS Plaza Bali
			Jl. Sunset Road No.68, Kuta, Kec. Kuta, Kabupaten Badung, Bali 80361, Kuta, Badung Kabupaten, Bali, Indonesia
			Jl. Sunset Road No.77B, Kuta, Kec. Kuta, Kabupaten Badung, Bali 80361Eccos Living Plaza`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"sport", "marathon", "run", "health", "maraton"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1705635876-2Ur1wEgiMGwupwRvIsF7mIje7owi6oTL.jpeg?width=1024&quality=90", true),
			Latitude:      null.NewString("-8.71603", true),
			Longitude:     null.NewString("115.18497", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Sports", "Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 9, 19, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 10, 4, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("KEMBALI BERKARYA 2024 : GUYON WATON", true),
			Description: null.NewString(`
			RAWS CHARITY RUN 2024 yang akan diselenggarakan pada 18 Februari 2024 di ECCOS Living Plaza Bali dengan konsep kegiatan Charity Fun Run 11K, bertujuan untuk mempromosikan wisata Bali dalam konsep wisata Kesehatan dengan menerapkan konsep health, wellness dan lifestyle. Event ini akan dilaksanakan di Bali peserta terdiri dari pengusaha dan masyarakat umum. Kegiatan RAWS CHARITY RUN 2024 ini diadakan bersamaan dengan Bali Health & Sport Tourism Expo atau BHESTE yang akan dilaksanakan pada tanggal 16-18 Februari 2024 di Eccos Living Plaza Bali.
			`, true),
			Location: null.NewString(`GOR Satria (GOR Satria Kota Purwokerto)
			Jl. Prof. Dr. Suharso No.8, Mangunjaya, Purwokerto Lor, Kec. Purwokerto Tim., Kabupaten Banyumas, Ja, Indonesia, Kabupaten Banyumas, Purwokerto`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"sport", "marathon", "run", "health", "maraton"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1706933162-X5G1Ib99JiI7UQgF1tjpKkLFUgV45a4y.jpg?width=1024&quality=90", true),
			Latitude:      null.NewString("-7.41584", true),
			Longitude:     null.NewString("109.25067", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 4, 23, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 4, 24, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 4, 26, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 4, 27, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Nextopia Festival Pool Party 2.0", true),
			Description: null.NewString(`
			Nextopia adalah sebuah event thematic yang akan memberikan wholesome experience kepada para pengunjung.

Lokasi : Hotel Sultan Jakarta

Tanggal : 26 - 27 April 2024

Jam : 20.00 - 02.00 (dini hari)

Dimeriahkan oleh 24 DJ & 2 MC

Jangan beli tiketnya, beli experience nya
			`, true),
			Location: null.NewString(`The Sultan Hotel Jl. Jend. Gatot Subroto, Gelora, Tanah Abang, Jakarta Pusat, Jakarta, Indonesia
Pool & Grass Area`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"poolparty", "party", "DJ", "hotelsultan", "Jakarta"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1708696574-Etyb9QHYbQbYBqEvH6WZQgRysruxVGOK.jpeg?width=1024&quality=90", true),
			Latitude:      null.NewString("-6.21852", true),
			Longitude:     null.NewString("106.81020", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Entertainment & Performance", "Expo", "Competition", "Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 5, 23, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 5, 24, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 5, 26, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 5, 27, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("ELFEST 2024", true),
			Description: null.NewString(`
		ELFEST Merupakan sebuah Festival Musik yang dipersembahkan oleh Politeknik TEDC Bandung dan dipromotori oleh sekelompok anak muda dari berbagai kalangan mahasiswa dan mahasiswi di indonesia. 
Selain festival musik ELFEST juga menghadirkan beberapa kegiatan yang menunjang kemajuan bangsa indonesia.
			`, true),
			Location:      null.NewString(`Cimahi, Kecamatan Cimahi Tengah, Cimahi Kota, Jawa Barat, Indonesia`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"festivalmusik", "robotik", "Bandung", "musik", "elfest"},
			Image:         null.NewString("https://d3hzxw6fwm263k.cloudfront.net/uploads/events/photos/1702985782-XBx8gkigpSDHWfFsd8E1R6M3etkiyClj.jpg?width=1024&quality=90", true),
			Latitude:      null.NewString("-6.88404", true),
			Longitude:     null.NewString("107.54104", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 23, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 24, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 6, 29, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 6, 29, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("APAC Marketers Middy's Indonesia", true),
			Description: null.NewString(`Bergabunglah bersama kami dalam Acara APAC Marketers Middy's pada 29 Februari, di mana para pemasar dari seluruh wilayah dapat menjalin hubungan, berbagi wawasan, dan bersenang-senang.
APAC Marketers Middy's Jakarta - 29 Juni

Minuman ringan bersama (middy’s) setiap pertengahan bulan, pertengahan minggu, dan pertengahan malam di tengah kota.

Tanpa Pitches Penjualan
Tanpa Agenda
Stark Taproom, Jakarta Indonesia

Elysee Lot 21 SCBD Lantai 1, RT.7/RW.1,

Senayan, Jakarta, Kota Jakarta Selatan, Jakarta 12190, Indonesia

Acara ini diselenggarakan oleh Kaliber Asia - anggota berdedikasi dari komunitas APAC Marketers.

Bersama-sama dengan Komunitas Penggemar Pemasaran - komunitas pemasaran terkemuka di Indonesia, dan Apiary Academy.

Tak Tertapis: Menavigasi 2024

#penghematananggaran #roi #brandvsperformance #tiktok

Glenn Karela, CPM (Asia) - Product Marketing Manager di KFC Indonesia

Rio Raditya - COO of Salvo

Wing Firmanperkasa - CEO of ANAKBRAND, Mantan Kepala Pemasaran di Gojek

Direktur Pemasaran Senior Coca Cola

Dipandu oleh Robert Lai - CEO Kaliber dan Kepala Komunitas APAC Marketers

Bergabunglah dalam acara fantastis di mana para pemasar APAC berkumpul untuk bersosialisasi, berbagi ide, dan bersenang-senang! Pertemuan langsung kami akan diadakan di Stark Taproom, Jakarta - tempat yang trendi dengan atmosfer sempurna untuk berjejaring dan menikmati minuman lezat.

Terhubunglah dengan sesama pemasar, tukar wawasan industri, dan bangun hubungan berharga.

Jangan lewatkan kesempatan menarik ini untuk memperluas jaringan profesional Anda dan bersenang-senang di tengah individu sejiwa.

Catat tanggal 29 Juni dan bersiaplah untuk malam yang tak terlupakan di APAC Marketers Middy's!

Tanpa biaya tiket

Letakkan tas di area yang ditentukan

Jika Anda melihat seseorang berdiri sendirian, ajaklah berbicara

Jika ada seseorang yang ingin Anda temui - beri tahu seseorang

Terakhir, setelah bekerja, hal terakhir yang ingin didengar oleh siapa pun adalah pitch penjualan, jadi tolong jangan menjual apa pun

APAC Marketers adalah komunitas anggota yang terdiri dari lebih dari 1.600+ pemasar di seluruh wilayah. Kami memiliki grup slack pribadi yang eksklusif untuk anggota. 
Jika Anda ingin bergabung, kirim email ke Robert@kaliber.asia dan kami akan menilai kesesuaian Anda dengan komunitas.
			`, true),
			Location:      null.NewString(`STARK Taproom (Stark Beer)Elysee Lot 21 SCBD 1st Floor, RT.7/RW.1, Senayan, Jakarta, South Jakarta City Jakarta, Jakarta 12190`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"bisnis", "marketers", "APAC", "business", "SCBD"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F702725159%2F1983780875573%2F1%2Foriginal.20240222-074223?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C2160%2C1080&s=3bfaaba9a904367150c52f158c2f34c9", true),
			Latitude:      null.NewString("-6.22591", true),
			Longitude:     null.NewString("106.81309", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 4, 17, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 4, 19, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 4, 21, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 4, 21, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("7 Figure Facebook Advertising Masterclass", true),
			Description: null.NewString(`
		Apakah Anda siap untuk merevolusi pendekatan Anda terhadap periklanan Facebook dan membuka kunci kesuksesan tujuh angka? Bergabunglah dengan kami untuk "7-Figure Facebook Advertising Masterclass" eksklusif kami, di mana Anda akan menemukan strategi, taktik, dan rahasia internal yang digunakan oleh para ahli industri untuk menghasilkan hasil besar.

Dalam masterclass yang penuh daya ini, Anda akan belajar bagaimana memanfaatkan potensi penuh periklanan Facebook untuk mendapatkan prospek, penjualan, dan pendapatan untuk bisnis Anda. Baik Anda seorang pemasar berpengalaman atau baru memulai, masterclass ini dirancang untuk membekali Anda dengan pengetahuan dan alat yang diperlukan untuk mencapai kesuksesan tak tertandingi dalam dunia periklanan digital.

Berikut yang dapat Anda harapkan untuk dipelajari:

Teknik Targeting Lanjutan: Buka kekuatan opsi penargetan Facebook untuk mencapai audiens ideal Anda dengan presisi dan akurasi.
Kreatif Iklan Berkonversi Tinggi: Temukan rahasia membuat kreatif iklan yang menarik perhatian dan mendorong tindakan.
Strategi Optimasi Kampanye: Pelajari cara mengoptimalkan kampanye iklan Facebook Anda untuk ROI maksimal dan efisiensi.
Strategi Scaling: Jelajahi strategi teruji untuk meningkatkan skala kampanye iklan Facebook Anda untuk mencapai audiens baru dan mendorong pertumbuhan eksponensial.
Analisis dan Pengukuran: Kuasai seni analisis iklan Facebook untuk melacak kinerja, mengukur hasil, dan mengoptimalkan untuk kesuksesan.
Baik Anda ingin menghasilkan prospek, meningkatkan penjualan, atau memperluas keberadaan merek Anda, masterclass ini akan memberi Anda wawasan dan teknik yang Anda butuhkan untuk mencapai tujuan Anda. Jangan lewatkan kesempatan ini untuk meningkatkan permainan periklanan Facebook Anda ke level berikutnya dan membuka kunci kesuksesan tujuh angka.
			`, true),
			Location:      null.NewString(`Online`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"advertising", "periklanan", "facebook", "edukasi"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F704651629%2F2018876648243%2F1%2Foriginal.20240225-050726?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C940%2C470&s=39e83c0518df87be4f2f6f28508416d1", true),
			Latitude:      null.NewString("0.0", true),
			Longitude:     null.NewString("0.0", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(true, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 4, 17, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 4, 19, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 4, 21, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 4, 21, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Unlock Your Path to Financial Freedom with Our Short Term Rental Workshop", true),
			Description: null.NewString(`

Apakah Anda siap untuk mengubah hidup Anda dan melangkah dengan percaya diri ke dunia kewirausahaan?

Bergabunglah dengan kami untuk sebuah workshop transformasional selama 3 jam untuk mengungkap rahasia membangun portofolio real estat yang menguntungkan berfokus pada Sewa Jangka Pendek (Short-Term Rentals/STRs).

Mengapa Workshop Ini Merupakan Game-Changer untuk Anda:

Wawasan dari Para Ahli: Pelajari mengapa STRs adalah tiket emas Anda menuju kesuksesan real estat dari para pemimpin industri.
Strategi yang Dapat Dilaksanakan: Temukan cara mengamankan unit STR pertama Anda, bernegosiasi dengan pemilik rumah, dan mengelola properti Anda dengan mudah.
Penawaran Langsung Eksklusif: Hadiri secara langsung untuk mengakses penawaran yang akan mendorong perjalanan real estat Anda — hanya tersedia selama workshop.
Peluang Jaringan: Terhubung dengan individu sejenis yang memiliki minat serupa dalam real estat dan kewirausahaan.

Bonus Khusus untuk Peserta:

Akses Gratis ke Pengetahuan Kelas Dunia: Amankan tempat Anda tanpa biaya! Ini adalah kesempatan Anda untuk mendapatkan wawasan berharga tanpa biaya yang mahal.

Rekaman Workshop & Transkrip: Apakah Anda melewatkan sesuatu atau ingin mengulang poin-poin kunci? Dapatkan rekaman dan transkrip workshop dengan biaya tambahan kecil.

Tanpa Pemutaran Ulang, Tanpa Alasan:

Berkomitmenlah pada masa depan Anda. Tidak ada pemutaran ulang, tetapi Anda dapat menjaga momentum dengan opsi untuk membeli rekaman dan transkrip.
			`, true),
			Location:      null.NewString(`Online`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"workshop", "finansial", "keuangan", "edukasi", "financial", "kewirausahaan"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F700514699%2F389505355397%2F1%2Foriginal.20240219-203744?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C400%2C960%2C480&s=329de834b71506bf33974824bed22d2c", true),
			Latitude:      null.NewString("0.0", true),
			Longitude:     null.NewString("0.0", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(true, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career", "Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 17, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 19, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 6, 21, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 6, 21, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Indonesia Big Meet", true),
			Description: null.NewString(`
Datang dan pelajari lebih lanjut tentang peluang bisnis untuk bekerja sama dengan Departemen Urusan Luar Negeri dan Perdagangan Australia dalam Program Pengembangan Australia.

Acara Indonesia Big Meet ini akan menjadi kesempatan bagi kontraktor lokal di Indonesia untuk mendiskusikan kerjasama dengan pemasok besar dan DFAT.

Hadirlah untuk:

Pembaruan tentang Pengaturan Strategis Program Pengembangan Indonesia dan Pipa Pengadaan
Tips praktis tentang 'Cara berbisnis dengan DFAT'
Peluang berjejaring dengan Kontraktor Manajemen
Studi kasus dan presentasi dari kontraktor lokal
Pesan kunci dari Komunitas Kontraktor Pengembangan Internasional
Panel Tanya Jawab
Tanggal: Rabu, 6 Maret 2024

Tempat: Akan diumumkan.

Waktu: 9.30 pagi - 2.00 sore (Pendaftaran dimulai dari jam 9:00 pagi)
			`, true),
			Location: null.NewString(`The Ritz-Carlton Jakarta, Mega Kuningan
no.1 Jalan Doktor Ide Anak Agung Gde Agung #Kav.E.1.1 Kecamatan Setiabudi, Daerah Khusus Ibukota Jakarta 12950`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"bisnis", "kerjasama", "strategi bisnis"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F680492019%2F222533957336%2F1%2Foriginal.20240123-051756?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C52%2C1640%2C820&s=61506b5b79c546f741991625149996f9", true),
			Latitude:      null.NewString("-6.22871", true),
			Longitude:     null.NewString("106.82721", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 17, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 19, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 7, 21, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 7, 21, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Virtual Learning - What Goes Where in Microsoft 365?", true),
			Description: null.NewString(`
Selamat datang di acara Pembelajaran Virtual kami - Memulai dengan Microsoft 365! Baik Anda baru mengenal Microsoft 365 atau ingin meningkatkan keterampilan Anda, acara online ini sangat cocok untuk Anda. Bergabunglah dengan kami untuk sesi yang menyenangkan dan interaktif di mana Anda akan belajar semua fitur luar biasa yang ditawarkan oleh Microsoft 365. Dalam sesi ini, Anda akan belajar:

Masuk ke Microsoft 365
Berlalu-lalang di dalam Microsoft 365
Gambaran umum tentang aplikasi yang termasuk dalam lingkungan Microsoft 365
Pengenalan tentang OneDrive
Pengenalan tentang SharePoint
Pengenalan tentang Microsoft Teams
Yang paling penting, kami akan menjawab pertanyaan besar "Apa yang Pergi Ke Mana!" Jangan lewatkan kesempatan ini untuk meningkatkan kemampuan teknologi Anda dari kenyamanan rumah Anda sendiri. Sampai jumpa di sana!
			`, true),
			Location:      null.NewString(`Online`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"microsoft", "microsoft365", "virtuallearning", "online", "edukasi"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F700864209%2F222190426335%2F1%2Foriginal.20240220-082237?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C2160%2C1080&s=9c0ba8538fa3c73a9c612db7be0e587f", true),
			Latitude:      null.NewString("0.0", true),
			Longitude:     null.NewString("0.0", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(true, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 7, 3, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 7, 3, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("MEGABUILD Indonesia (MBI)", true),
			Description: null.NewString(`
MEGABUILD Indonesia (MBI) adalah acara pameran bisnis bahan bangunan, arsitektur, desain interior, dan konstruksi terkemuka di Indonesia yang diselenggarakan oleh industri untuk industri.

MBI akan menjadi pameran paling komprehensif tentang teknologi terkini, solusi, inovasi, produk, dan tren desain untuk industri konstruksi, interior, dan bangunan di Indonesia.

Ini adalah acara tahunan B2B2C yang menampilkan prominently di kalender perdagangan regional untuk komunitas global dari kontraktor, arsitek, desainer interior, profesional konstruksi, dan pemilik rumah.

MBI berusaha memberikan pengalaman berharga dan membuat peluang terhubung selama dan setelah hari-hari pameran!
			`, true),
			Location: null.NewString(`Balai Sidang Jakarta Convention Center

1 Jalan Gatot Subroto Kecamatan Tanah Abang, Daerah Khusus Ibukota Jakarta 10270`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"arsitek", "interiordesign", "bangunan", "konstruksi", "teknologi", "pameran"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F531789509%2F287155506737%2F1%2Foriginal.20230608-074410?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C270%2C1080%2C540&s=a4743101e9f444a8fe5666c6dc733f10", true),
			Latitude:      null.NewString("-6.21407", true),
			Longitude:     null.NewString("106.80724", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 5, 8, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 5, 9, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 5, 12, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("KERAMIKA Indonesia (KMI)", true),
			Description: null.NewString(`
KERAMIKA INDONESIA akan membawa produk dari produsen keramik dan pemasok bahan baku, peralatan, dan mesin, kepada salah satu KONSUMEN TERBESAR LANTAI KERAMIK & PERLENGKAPAN SANITER DI ASIA TENGGARA, INDONESIA.

MEMPERSEMBAHKAN 3 DUNIA DALAM SATU ATAP:

DUNIA MESIN & TEKNOLOGI - mencakup seluruh industri manufaktur keramik mencakup peralatan, teknologi, solusi, pasokan, dan finishing.

DUNIA KERAMIK - memamerkan peluncuran baru, desain, dan solusi dari Asosiasi Industri Keramik Indonesia (ASAKI), membawa yang terbaik di kelasnya ke dalam zona World of Ceramics eksklusif.

DUNIA BAHAN BAKU - menampilkan mineral inovasi bahan baku keramik.
			`, true),
			Location: null.NewString(`Balai Sidang Jakarta Convention Center

1 Jalan Gatot Subroto Kecamatan Tanah Abang, Daerah Khusus Ibukota Jakarta 10270`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"lantai", "keramik", "bangunan", "konstruksi", "teknologi", "pameran"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F531797039%2F287155506737%2F1%2Foriginal.20230608-080808?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C270%2C1080%2C540&s=1fd97eda2c18ce28a1e9450f0b20d699", true),
			Latitude:      null.NewString("-6.21407", true),
			Longitude:     null.NewString("106.80724", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 5, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 5, 8, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 5, 15, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 5, 15, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Unistart International Education Expo 2024", true),
			Description: null.NewString(`
Unistart International Education Expo 2024 - acara yang wajib dihadiri bagi para siswa yang mencari pengalaman pendidikan internasional dan ingin melanjutkan studi di luar negeri! Acara ini akan berlangsung pada hari Sabtu, 9 Maret 2024 di HARRIS Hotel & Conventions Kelapa Gading. Jelajahi opsi program, dan kumpulkan informasi berharga untuk membuat keputusan yang tepat tentang pendidikanmu.



Temukan berbagai universitas, perguruan tinggi, dan sekolah bahasa. Berkonsultasilah dengan perwakilan universitas, dapatkan semua informasi tentang program beasiswa yang tersedia. Kamu juga dapat menguji keterampilan bahasa Inggrismu dengan Uji Simulasi IELTS dan mengikuti seminarnya bersama The British Institute, serta belajar cara menulis Personal Statement untuk penerimaan Universitas bersama Ms. Trifty Qurrota Aini, Indonesia Country Manager UWE Bristol atau dibimbing menuju karir impianmu dengan cara memilih jurusan yang tepat bersama Ms. Ariyani Mawardi, Head of USG Education Kelapa Gading Campus.

Entah kamu tertarik pada program sarjana, pasca sarjana, atau bahasa, expo ini memiliki solusi untuk semua orang.



Jangan lewatkan kesempatan luar biasa ini dan menangkan SATU Tiket Pulang-Pergi (PP) Qantas ke Sydney atau Melbourne !



Tandai kalendermu dan pastikan untuk hadir di Unistart International Education Expo 2024.
			`, true),
			Location:      null.NewString(`HARRIS Hotel & Conventions Kelapa Gading RT.13 Jalan Boulevard Barat Raya #Blok M Kec. Klp. Gading, Daerah Khusus Ibukota Jakarta 14240`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"expo", "pameran", "education", "edukasi", "beasiswa", "scholarship"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F696150759%2F1340129042443%2F1%2Foriginal.20240213-150952?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C2160%2C1080&s=f03bce6799600a9100ba30ef692a8fca", true),
			Latitude:      null.NewString("-6.15724", true),
			Longitude:     null.NewString("106.91016", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 6, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 6, 4, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 6, 4, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 6, 6, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("PCE Yinghe Oral Care Expo Jakarta Station", true),
			Description: null.NewString(`PCE Yinghe Personal Care Expo Jakarta Station 2024 membantu perusahaan-perusahaan China untuk berpartisipasi dengan lancar dalam pameran luar negeri dalam bentuk "negosiasi online domestik, pameran dagang offline luar negeri". Melalui layanan rantai penuh "barang pameran berangkat, pameran online peserta, pembeli di lokasi, dan negosiasi instan", dan siklus layanan ultra-panjang selama 31 hari untuk mencapai pencocokan bisnis selama 31 hari, menciptakan mode pameran baru yang "berkualitas tinggi dan efisien".

PCE Yinghe Personal Care Expo Jakarta Station 2024 akan berlangsung dari 30 Mei hingga 1 Juni 2024 di Jakarta International Expo, pameran mencakup kekuatan utama baru di bidang perawatan pribadi seperti rantai pasokan, pemilik merek, dan sumber daya ekologis, dengan pameran yang lebih kaya dan kategori yang lebih lengkap. Ini menghubungkan pembeli profesional luar negeri dengan merek, produsen berkualitas tinggi, dan perusahaan ekspor dari China, menampilkan tren pengembangan terbaru dalam industri produk perawatan pribadi, memperkenalkan produk-produk baru utama, meningkatkan pangsa pasar produk perusahaan China, dan bertemu mitra bisnis baru. Pameran ini bertujuan untuk membangun platform layanan berkualitas tinggi untuk industri produk perawatan pribadi yang mengintegrasikan pameran, kerja sama perdagangan, peluncuran produk baru, pembelajaran, dan komunikasi. Ini mewujudkan integrasi multi-saluran dan seluruh rantai industri, mendapatkan informasi pengadaan terbaru dan peluang tender, dan mencari pengembangan industri yang baik.
			`, true),
			Location:      null.NewString(`Jakarta International Expo Arena JIExpo Kemayoran Jakarta, 10620`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"personalcare", "skincare", "produkrambut", "rambut", "expo", "china"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F633349789%2F300539367512%2F1%2Foriginal.20231102-025427?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=1%2C22%2C4320%2C2160&s=42fbb1d2dd9f1269dbab109fa79936de", true),
			Latitude:      null.NewString("-6.14756", true),
			Longitude:     null.NewString("106.84608", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Expo"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 4, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 4, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 6, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("IFINEXPO Jakarta 2024", true),
			Description: null.NewString(`IFINEXPO Jakarta 2024 merupakan platform besar yang 
mengumpulkan sumber daya berkualitas tinggi dari seluruh dunia, menggali peluang bisnis yang tak terbatas, 
dan mempromosikan kerjasama internasional. Acara ini menawarkan berbagai keuntungan, termasuk peluang pertumbuhan 
bisnis, kesempatan untuk mempromosikan merek, dan platform untuk berkolaborasi dan bertukar ide dengan pemimpin industri. 
Dengan pasar keuangan Indonesia yang berkembang pesat dan potensial besar, IFINEXPO Jakarta 2024 adalah acara yang 
tidak boleh Anda lewatkan untuk meraih kesuksesan di pasar keuangan Indonesia.
			`, true),
			Location:      null.NewString(`Jalan Gatot Subroto Kecamatan Tanah Abang, Daerah Khusus Ibukota Jakarta 10270`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"bisnis", "pasar uang", "peluang bisnis", "jakarta", "expo"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F658561599%2F700916941403%2F1%2Foriginal.20231214-115938?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C221%2C1280%2C640&s=84e0faa3c06d4e4c3375f5be61445a17", true),
			Latitude:      null.NewString("-6.22515", true),
			Longitude:     null.NewString("106.81675", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 4, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 4, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 6, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("The MarTech Summit Jakarta June 2024", true),
			Description: null.NewString(`The MarTech Summit akan datang ke Jakarta untuk kedua kalinya, 
menghadirkan beberapa pemikir terbaik dan perusahaan paling inovatif di industri ini. Dengan diskusi panel 
yang penuh wawasan, obrolan santai, presentasi kunci yang memprovokasi pemikiran, dan sesi bulat interaktif, 
Anda akan meninggalkan acara ini dilengkapi dengan alat dan strategi terbaru untuk meningkatkan upaya pemasaran Anda. 
Acara ini akan memberikan platform unik bagi para CMO, Kepala, Direktur, dan lainnya untuk berjejaring, belajar, 
dan tetap berada di depan dalam dunia MarTech yang selalu berubah.
Acara ini dirancang untuk Eksekutif Tingkat Senior dalam fungsi-fungsi seperti:
Pemasaran & Teknologi | Pengalaman Pelanggan (CX) & Keterlibatan | Kesetiaan Merek & Retensi | Data & wawasan Konsumen | Pemasaran E-commerce | Strategi Digital | OmniChannel | Inovasi | Media Sosial | Strategi Konten & Perenungan | CRM | Pemasaran Produk | Otomatisasi | Transformasi Digital & Pertumbuhan | Pengalaman Digital | Teknologi Informasi | SEO & SEM
			`, true),
			Location:      null.NewString(`Jakarta City Centre Jakarta, 10310`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"bisnis", "industri", "inovasi", "jakarta", "edukasi"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F666577639%2F308052947969%2F1%2Foriginal.20240103-110837?w=940&auto=format%2Ccompress&q=75&sharp=10&s=224fe31044b80bb5beb46311f5134904", true),
			Latitude:      null.NewString("-6.22127", true),
			Longitude:     null.NewString("106.80914", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 5, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("Transforming businesses and unlocking innovation with SoftwareOne", true),
			Description: null.NewString(`Banyak organisasi mengatakan bahwa laju transformasi digital 
dalam industri mereka telah meningkat selama beberapa tahun terakhir. Seiring teknologi cloud 
memperluas akses terhadap inovasi di luar departemen IT, kami ingin membantu pelanggan mempercepat 
pertumbuhan bisnis mereka dengan memodernisasi aplikasi, infrastruktur, dan data mereka dengan SoftwareOne dan Microsoft.
			`, true),
			Location:      null.NewString(`The Westin Jakarta Kav.C-22A Jalan Haji R. Rasuna Said Kecamatan Setiabudi, Daerah Khusus Ibukota Jakarta 12940`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"industri", "IT", "modernisasi", "jakarta", "edukasi", "microsoft"},
			Image:         null.NewString("https://img.evbuc.com/https%3A%2F%2Fcdn.evbuc.com%2Fimages%2F694960229%2F455165111728%2F1%2Foriginal.20240212-024629?w=940&auto=format%2Ccompress&q=75&sharp=10&rect=0%2C0%2C2160%2C1080&s=98c6242376eba7bb3fd0829052cac942", true),
			Latitude:      null.NewString("-6.22355", true),
			Longitude:     null.NewString("106.83384", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},

		{
			EventcreatorID:  null.NewInt(3, true),
			Category:        []string{"Education & Career"},
			PreregisterDate: null.NewTime(time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC), true),
			EndregisterDate: null.NewTime(time.Date(2024, 8, 5, 0, 0, 0, 0, time.UTC), true),
			StarteventDate:  null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			EndeventDate:    null.NewTime(time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC), true),
			Title:           null.NewString("SEMINAR NASIONAL Pharmacy Informatics", true),
			Description: null.NewString(`✨ HIMAFA UNIMMA MEMPERSEMBAHKAN ✨
Hai, sobat farma!

HIMAFA UNIMMA mengadakan Seminar Nasional nih, Pada hari Sabtu, 02 Maret 2024 (08.00-Selesai) baik secara offline maupun online.



💫 SEMINAR NASIONAL 💫

Dengan Tema "Pharmacy Informatics : Developing pharmaceutical service and drug discovery and their use in the era of the industrial revolution 4.0"

 

Tentunya Sobat Farma yang ikut berpartisipasi akan mendapatkan Sertifikat ber-SKP IAI dan PAFI .

Tunggu apa lagi sobat? Yok segera daftar 🫵🏻



🗣️ Pemateri :

1. Apt. Lalu Muhammad Irham M.Farm., Ph.D

2. Apt.Nadia Saptarina M.Pharm

3. Arief Kusuma Wardani M.Pharm.S.ci



👤 Moderator :

Apt. Herma Fanani Agusta, M.Sc



 OFFLINE 

🏫 Lokasi :

Auditorium Kampus 1 Unimma

Jl. Tidar No.21, Magersari, Kec. Magelang Sel., Kota Magelang, Jawa Tengah 59214
			`, true),
			Location:      null.NewString(`Auditorium Kampus 1 Unimma, Kota Magelang`, true),
			AverageRating: null.NewFloat64(0, true),
			IsFinished:    null.NewBool(false, true),
			Tags:          []string{"industri", "farmasi", "informatika", "jakarta", "edukasi", "seminar"},
			Image:         null.NewString("https://eventkampus.com/data/event/poster/22/_thumbnail/600x600/4493-seminar-nasional-pharmacy-informatics.jpeg", true),
			Latitude:      null.NewString("-7.48820", true),
			Longitude:     null.NewString("110.21886", true),
			GuestStar:     []string{},
			IsOnline:      null.NewBool(false, true),
		},
	}

	return EventsToInsert
}
