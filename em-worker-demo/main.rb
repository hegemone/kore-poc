# Very simple, vanilla eventmachine based demo for illustrating reactor

require 'eventmachine'

EM.threadpool_size = 32
MAX_JOBS = 64
MAX_TICK = 10
MAX_SLEEP = 3


def main
  puts '============================================================'
  puts '                       EM Reactor Demo'
  puts '============================================================'
  Engine.new().start
end

class DerpJob
  attr_reader :name

  def initialize(id, tick, done)
    @id = id
    @tick = tick
    @done = done
    @name = "foo_#{@id}"
    @total_ticks = Random.rand((1..MAX_TICK+1))
    @sleep = Random.rand((1..MAX_SLEEP+1))
  end

  def start
    EM.defer { self.do_work } # Crux of the async work
  end

  def do_work
    (0..@total_ticks).each do |i|
      self.send_tick(i)
      sleep @sleep
    end

    @done.call(msg("Aw yiss! All done."))
  end

  def send_tick(tick_num)
    p_str = if tick_num == 0 then "0.0%" else
              sprintf("%.1f%%", tick_num.to_f / @total_ticks * 100)
            end
    @tick.call(self.msg("progress: #{p_str}"))
  end

  def msg(val)
    { job: self, val: val }
  end
end

class Engine
  def start
    puts 'Engine#start'

    EM.run do
      puts "Spawning #{MAX_JOBS} job..."

      # Execute `EM.threadpool_size` jobs in parallel
      (1..MAX_JOBS).each do |id|
        DerpJob.new(id, method(:on_tick), method(:on_done)).start
      end

      # Prove job work is non-blocking, evented, and running in parallel
      puts "Started all the work!"
    end
  end

  def on_tick(msg)
    print "Engine#on_tick: job[#{msg[:job].name}], val:[#{msg[:val]}]\n"
  end

  def on_done(msg)
    print "Engine#on_done: job[#{msg[:job].name}], val:[#{msg[:val]}]\n"
  end
end

main
