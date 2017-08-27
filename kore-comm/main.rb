require 'logging'
require 'eventmachine'
require 'dry-auto_inject'
require 'dry-container'
require 'pry'

#############################################################
# Deps setup and injection
#############################################################
# Full blown project will use a proper config setup
# Not relevant for demo purposes, so we'll use a map
_config = {
  foo: 'bar',
  em_threadpool_size: 20 # 20 is default, just showing it's configurable
}
_log = Logging.logger(STDOUT)

container = Dry::Container.new
container.register(:log, _log)
container.register(:config, _config)
Inject = Dry::AutoInject(container)

EM.threadpool_size = _config[:em_threadpool_size]

#############################################################
# Fire up Kore
#############################################################
Dir.glob('./lib/**/*.rb') { |file| require file }
Kore::Machinery::Engine.new().start
