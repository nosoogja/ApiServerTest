using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace ApiServer.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class CatchmeController : ControllerBase
    {
        private static int checkNum = 0;

        private readonly ILogger<WeatherForecastController> _logger;

        //public CatchmeController(ILogger<WeatherForecastController> logger)
        public CatchmeController()
        {
            //_logger = logger;
        }

        [HttpGet]
        public string Get()
        {
                return $"{checkNum++}";
        }
    }
}
