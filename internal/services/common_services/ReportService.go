package common_services

type ReportService struct {
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (r *ReportService) ExportToExcel(data interface{}, lang string) (error, []byte) {

}
