
function Home() {
    React.useEffect(() => {
        new Swiper(".mySwiper", {
            slidesPerView: 1,
            spaceBetween: 20,
            loop: true,
            breakpoints: {
                640: { slidesPerView: 2 },
                1024: { slidesPerView: 3 },
            },
            pagination: { el: ".swiper-pagination", clickable: true },
            navigation: { nextEl: ".swiper-button-next", prevEl: ".swiper-button-prev" },
        });
    }, []);

    return (
        <div className="p-6">
            <h1 className="text-2xl font-bold mb-4">Halaman Home</h1>
            <div className="swiper mySwiper">
                <div className="swiper-wrapper">
                    <div className="swiper-slide">
                        <div className="bg-white p-4 rounded-xl shadow text-center">
                            <img src="https://picsum.photos/200/150?random=1" className="mx-auto mb-2 rounded-lg" />
                            <p>Card 1</p>
                        </div>
                    </div>
                    <div className="swiper-slide">
                        <div className="bg-white p-4 rounded-xl shadow text-center">
                            <img src="https://picsum.photos/200/150?random=2" className="mx-auto mb-2 rounded-lg" />
                            <p>Card 2</p>
                        </div>
                    </div>
                    <div className="swiper-slide">
                        <div className="bg-white p-4 rounded-xl shadow text-center">
                            <img src="https://picsum.photos/200/150?random=3" className="mx-auto mb-2 rounded-lg" />
                            <p>Card 3</p>
                        </div>
                    </div>
                    <div className="swiper-slide">
                        <div className="bg-white p-4 rounded-xl shadow text-center">
                            <img src="https://picsum.photos/200/150?random=4" className="mx-auto mb-2 rounded-lg" />
                            <p>Card 4</p>
                        </div>
                    </div>
                </div>

                {/* Pagination & Navigation */}
                <div className="swiper-pagination"></div>
                <div className="swiper-button-next"></div>
                <div className="swiper-button-prev"></div>
            </div>
        </div>
    );
}
