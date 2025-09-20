const { useState, useEffect } = React;

function Detail() {
    document.title = "detail"
    const [berita, setBerita] = useState([]);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(1);
    const [searchTerm, setSearchTerm] = useState(""); 
    const limit = 5;  

    useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await fetch("http://localhost:6969/user/index");
                const data = await res.json();
                setBerita(data.data || []);
            } catch (err) {
                console.error("Gagal fetch data:", err);
            } finally {
                setLoading(false);
            }
        };
        fetchData();
    }, []);

    const filteredData = berita.filter(
        (item) =>
            item.username.toLowerCase().includes(searchTerm.toLowerCase()) ||
            item.email.toLowerCase().includes(searchTerm.toLowerCase())
    );

    // Pagination
    const startIndex = (page - 1) * limit;
    const currentData = filteredData.slice(startIndex, startIndex + limit);
    const totalPages = Math.ceil(filteredData.length / limit);

    return (
        <div className="p-6">
            <h1 className="text-2xl font-bold mb-4">Halaman Detail</h1>
            <p className="mb-6 text-gray-600">Ini adalah halaman detail dengan tabel user + search.</p>

            {/* 🔍 Input Search */}
            <div className="mb-4">
                <input
                    type="text"
                    placeholder="Cari username atau email..."
                    value={searchTerm}
                    onChange={(e) => {
                        setSearchTerm(e.target.value);
                        setPage(1); // reset ke halaman pertama saat search
                    }}
                    className="border px-3 py-2 rounded w-full md:w-1/3 focus:outline-none focus:ring focus:ring-blue-300"
                />
            </div>

            {loading ? (
                <p className="text-center">Loading...</p>
            ) : (
                <>
                    <table className="table-auto w-full border-collapse border border-gray-300 shadow-md rounded-lg">
                        <thead className="bg-gray-100">
                            <tr>
                                <th className="border border-gray-300 px-4 py-2">No</th>
                                <th className="border border-gray-300 px-4 py-2">Username</th>
                                <th className="border border-gray-300 px-4 py-2">Email</th>
                            </tr>
                        </thead>
                        <tbody>
                            {currentData.length > 0 ? (
                                currentData.map((item, index) => (
                                    <tr key={item.id} className="hover:bg-gray-50">
                                        <td className="border border-gray-300 px-4 py-2 text-center">
                                            {startIndex + index + 1}
                                        </td>
                                        <td className="border border-gray-300 px-4 py-2">{item.username}</td>
                                        <td className="border border-gray-300 px-4 py-2">{item.email}</td>
                                    </tr>
                                ))
                            ) : (
                                <tr>
                                    <td colSpan="3" className="text-center py-4 text-gray-500">
                                        Tidak ada data ditemukan
                                    </td>
                                </tr>
                            )}
                        </tbody>
                    </table>

                    {/* Pagination */}
                    {filteredData.length > 0 && (
                        <div className="flex justify-center mt-4 space-x-2">
                            <button
                                onClick={() => setPage((p) => Math.max(p - 1, 1))}
                                className="px-3 py-1 bg-gray-200 rounded disabled:opacity-50"
                                disabled={page === 1}
                            >
                                Prev
                            </button>
                            <span className="px-3 py-1">
                                {page} / {totalPages}
                            </span>
                            <button
                                onClick={() => setPage((p) => Math.min(p + 1, totalPages))}
                                className="px-3 py-1 bg-gray-200 rounded disabled:opacity-50"
                                disabled={page === totalPages}
                            >
                                Next
                            </button>
                        </div>
                    )}
                </>
            )}
        </div>
    );
}
